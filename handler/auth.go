package handler

import (
	"context"
	"errors"

	"github.com/micro/go-micro/v2/util/log"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/lecex/user/proto/auth"
	userPb "github.com/lecex/user/proto/user"

	"github.com/lecex/user/service"
	"github.com/lecex/user/service/repository"
)

// Auth 授权服务处理
type Auth struct {
	TokenService service.Authable
	Repo         repository.User
}

// AuthById 只通过 id 获取 jwt token
// bug 只能服务之间调用如果前端调用会不验证获取权限
func (srv *Auth) AuthById(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(&userPb.User{
		Id:     req.User.Id,
		Mobile: req.User.Mobile,
	})
	if err != nil {
		log.Log(err)
		return err
	}
	if user != nil {
		req.User.Id = user.Id
		t, err := srv.TokenService.Encode(req.User)
		if err != nil {
			log.Log(err)
			return err
		}
		res.Token = t

		res.User = &pb.User{
			Id:       user.Id,
			Username: user.Username,
			Mobile:   user.Mobile,
			Email:    user.Email,
			Name:     user.Name,
			Avatar:   user.Avatar,
			Origin:   user.Origin,
		}
	}
	return nil
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 在 part3 中直接传参 &pb.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// 即 req.Password 与 u.Password 都是加密后的密码
	// 将无法通过验证
	user, err := srv.Repo.Get(&userPb.User{
		Id:       req.User.Id,
		Username: req.User.Username,
		Email:    req.User.Email,
		Mobile:   req.User.Mobile,
	})
	if err != nil {
		return errors.New("用户不存在")
	}
	if user != nil {
		// 进行密码验证
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.User.Password)); err != nil {
			log.Log(err)
			return errors.New("密码错误")
		}
		req.User.Id = user.Id
		req.User.Username = user.Username
		req.User.Email = user.Email
		req.User.Mobile = user.Mobile
		t, err := srv.TokenService.Encode(req.User)
		if err != nil {
			log.Log(err)
			return err
		}
		res.Token = t
		res.User = &pb.User{
			Id:       user.Id,
			Username: user.Username,
			Mobile:   user.Mobile,
			Email:    user.Email,
			Name:     user.Name,
			Avatar:   user.Avatar,
			Origin:   user.Origin,
		}
	}
	return nil
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// Decode token
	if req.Token == "" {
		return errors.New("请输传入 Token")
	}
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		log.Log(err)
		return err
	}
	if claims.User.Id == "" {
		return errors.New("无效用户")
	}
	res.User = claims.User
	res.Valid = true
	return err
}

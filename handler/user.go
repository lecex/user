package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/user/proto/user"
	"github.com/lecex/user/service/repository"

	eventPB "github.com/lecex/core/proto/event"
	"golang.org/x/crypto/bcrypt"
)

// User 用户结构
type User struct {
	Repo      repository.User
	Publisher micro.Publisher
}

// Exist 用户是否存在
func (srv *User) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Repo.Exist(req.User)
	return err
}

// List 获取所有用户
func (srv *User) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	users, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		log.Log(err)
		return err
	}
	res.Total = total
	res.Users = users
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(req.User)
	if err != nil {
		log.Log(err)
		return err
	}
	res.User = user
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.User.Password = string(hashedPass)
	user, err := srv.Repo.Create(req.User)
	if err != nil {
		log.Log(err)
		res.Valid = false
		return fmt.Errorf("创建用户失败: %s", err)
	}
	res.User = user
	res.Valid = true
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 密码不为空时才可以修改密码
	if req.User.Password != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Log(err)
			return err
		}
		req.User.Password = string(hashedPass)
	}
	valid, err := srv.Repo.Update(req.User)
	if err != nil {
		log.Log(err)
		res.Valid = false
		return fmt.Errorf("更新用户失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除用户用户
func (srv *User) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.User)
	if err != nil {
		log.Log(err)
		res.Valid = false
		return fmt.Errorf("删除用户失败")
	}
	if valid {
		if err := srv.publish(ctx, req.User, "user_delete"); err != nil {
			return err
		}
	}
	res.Valid = valid
	return err
}

// publish 消息发布
func (srv *User) publish(ctx context.Context, user *pb.User, topic string) (err error) {
	if srv.Publisher != nil {
		u, err := srv.Repo.Get(user)
		if err != nil {
			return err
		}
		data, _ := json.Marshal(&u)
		event := &eventPB.Event{
			UserId:     "",
			DeviceInfo: "",
			GroupId:    "",
			Topic:      topic,
			Data:       data,
		}
		return srv.Publisher.Publish(ctx, event)
	}
	return
}

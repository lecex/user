package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/user/proto/secretKey"

	"github.com/lecex/user/service/repository"
)

// SecretKey 秘钥结构
type SecretKey struct {
	Repo repository.SecretKey
}

// All 获取所有权限
func (srv *SecretKey) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	secretKeys, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.SecretKeys = secretKeys
	return err
}

// List 获取所有秘钥
func (srv *SecretKey) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	secretKeys, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.SecretKeys = secretKeys
	res.Total = total
	return err
}

// Get 获取秘钥
func (srv *SecretKey) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	secretKey, err := srv.Repo.Get(req.SecretKey)
	if err != nil {
		return err
	}
	res.SecretKey = secretKey
	return err
}

// Create 创建秘钥
func (srv *SecretKey) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.SecretKey)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加秘钥失败")
	}
	res.Valid = true
	return err
}

// Update 更新秘钥
func (srv *SecretKey) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.SecretKey)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新秘钥失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除秘钥
func (srv *SecretKey) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.SecretKey)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除秘钥失败")
	}
	res.Valid = valid
	return err
}

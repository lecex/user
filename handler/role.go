package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/user/proto/role"

	"github.com/lecex/user/service/repository"
)

// Role 角色结构
type Role struct {
	Repo repository.Role
}

// All 获取所有权限
func (srv *Role) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	roles, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.Roles = roles
	return err
}

// List 获取所有角色
func (srv *Role) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	roles, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Roles = roles
	res.Total = total
	return err
}

// Get 获取角色
func (srv *Role) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	role, err := srv.Repo.Get(req.Role)
	if err != nil {
		return err
	}
	res.Role = role
	return err
}

// Create 创建角色
func (srv *Role) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加角色失败")
	}
	res.Valid = true
	return err
}

// Update 更新角色
func (srv *Role) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新角色失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除角色
func (srv *Role) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Role)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除角色失败")
	}
	res.Valid = valid
	return err
}

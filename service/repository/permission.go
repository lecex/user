package repository

import (
	"fmt"
	// 公共引入
	"github.com/jinzhu/gorm"
	"github.com/lecex/core/util"
	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/user/proto/permission"
)

//Permission 仓库接口
type Permission interface {
	Create(permission *pb.Permission) (*pb.Permission, error)
	Delete(permission *pb.Permission) (bool, error)
	Update(permission *pb.Permission) (bool, error)
	Get(permission *pb.Permission) (*pb.Permission, error)
	All(req *pb.Request) ([]*pb.Permission, error)
	List(req *pb.ListQuery) ([]*pb.Permission, error)
	Total(req *pb.ListQuery) (int64, error)
}

// PermissionRepository 权限仓库
type PermissionRepository struct {
	DB *gorm.DB
}

// All 获取所有角色信息
func (repo *PermissionRepository) All(req *pb.Request) (permissions []*pb.Permission, err error) {
	if err := repo.DB.Find(&permissions).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return permissions, nil
}

// List 获取所有权限信息
func (repo *PermissionRepository) List(req *pb.ListQuery) (permissions []*pb.Permission, err error) {
	db := repo.DB
	limit, offset := util.Page(req.Limit, req.Page) // 分页
	sort := util.Sort(req.Sort)                     // 排序 默认 created_at desc
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&permissions).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return permissions, nil
}

// Total 获取所有权限查询总量
func (repo *PermissionRepository) Total(req *pb.ListQuery) (total int64, err error) {
	permissions := []pb.Permission{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&permissions).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取权限信息
func (repo *PermissionRepository) Get(permission *pb.Permission) (*pb.Permission, error) {
	if err := repo.DB.Where(&permission).Find(&permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

// Create 创建权限
// bug 无权限名创建权限可能引起 bug
func (repo *PermissionRepository) Create(p *pb.Permission) (*pb.Permission, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return p, fmt.Errorf("添加权限失败")
	}
	return p, nil
}

// Update 更新权限
func (repo *PermissionRepository) Update(p *pb.Permission) (bool, error) {
	if p.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	err := repo.DB.Model(&p).Where("id = ?", p.Id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除权限
func (repo *PermissionRepository) Delete(p *pb.Permission) (bool, error) {
	if p.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	err := repo.DB.Model(&p).Where("id = ?", p.Id).Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

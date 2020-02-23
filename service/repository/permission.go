package repository

import (
	"fmt"
	// 公共引入
	"github.com/jinzhu/gorm"
	"github.com/lecex/core/uitl"
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
	List(listQuery *pb.ListQuery, per *pb.Permission) ([]*pb.Permission, error)
	Total(req *pb.Permission) (int64, error)
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
func (repo *PermissionRepository) List(listQuery *pb.ListQuery) (permissions []*pb.Permission, err error) {
	db := repo.DB
	// 计算分页
	limit, offset := uitl.Page(listQuery.Limit, listQuery.Page)
	// 排序
	var sort string
	if listQuery.Sort != "" {
		sort = listQuery.Sort
	} else {
		sort = "id desc"
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&permissions).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return permissions, nil
}

// Total 获取所有权限查询总量
func (repo *PermissionRepository) Total(req *pb.Permission) (total int64, err error) {
	permissions := []pb.Permission{}
	db := repo.DB
	if err := db.Find(&permissions).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取权限信息
func (repo *PermissionRepository) Get(p *pb.Permission) (*pb.Permission, error) {
	if p.Id > 0 {
		if err := repo.DB.Model(&p).Where("id = ?", p.Id).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	if p.Service != "" && p.Method != "" {
		if err := repo.DB.Model(&p).Where("service = ?", p.Service).Where("method = ?", p.Method).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	if p.Name != "" {
		if err := repo.DB.Model(&p).Where("name = ?", p.Name).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	return p, nil
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
	id := &pb.Permission{
		Id: p.Id,
	}
	err := repo.DB.Model(id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除权限
func (repo *PermissionRepository) Delete(p *pb.Permission) (bool, error) {
	err := repo.DB.Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

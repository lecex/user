package service

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/user/proto/role"

	"github.com/jinzhu/gorm"
)

//RRepository 仓库接口
type RRepository interface {
	Create(role *pb.Role) (*pb.Role, error)
	Delete(role *pb.Role) (bool, error)
	Update(role *pb.Role) (bool, error)
	Get(role *pb.Role) (*pb.Role, error)
	All(req *pb.Request) ([]*pb.Role, error)
	List(req *pb.ListQuery) ([]*pb.Role, error)
	Total(req *pb.ListQuery) (int64, error)
}

// RoleRepository 角色仓库
type RoleRepository struct {
	DB *gorm.DB
}

// All 获取所有角色信息
func (repo *RoleRepository) All(req *pb.Request) (roles []*pb.Role, err error) {
	if err := repo.DB.Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

// List 获取所有角色信息
func (repo *RoleRepository) List(req *pb.ListQuery) (roles []*pb.Role, err error) {
	db := repo.DB
	// 分页
	var limit, offset int64
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = -1
	}

	// 排序
	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "id desc"
	}
	// 查询条件
	if req.Label != "" {
		db = db.Where("label like ?", "%"+req.Label+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

// Total 获取所有角色查询总量
func (repo *RoleRepository) Total(req *pb.ListQuery) (total int64, err error) {
	roles := []pb.Role{}
	db := repo.DB
	// 查询条件
	if req.Label != "" {
		db = db.Where("label like ?", "%"+req.Label+"%")
	}
	if err := db.Find(&roles).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取角色信息
func (repo *RoleRepository) Get(r *pb.Role) (*pb.Role, error) {
	if r.Id > 0 {
		if err := repo.DB.Model(&r).Where("id = ?", r.Id).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	if r.Label != "" {
		if err := repo.DB.Model(&r).Where("label = ?", r.Label).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	if r.Name != "" {
		if err := repo.DB.Model(&r).Where("name = ?", r.Name).Find(&r).Error; err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Create 创建角色
// bug 无角色名创建角色可能引起 bug
func (repo *RoleRepository) Create(r *pb.Role) (*pb.Role, error) {
	err := repo.DB.Create(r).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return r, fmt.Errorf("添加角色失败")
	}
	return r, nil
}

// Update 更新角色
func (repo *RoleRepository) Update(r *pb.Role) (bool, error) {
	if r.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Role{
		Id: r.Id,
	}
	err := repo.DB.Model(id).Updates(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除角色
func (repo *RoleRepository) Delete(r *pb.Role) (bool, error) {
	err := repo.DB.Delete(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

package repository

import (
	"fmt"
	// 公共引入
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/lecex/core/uitl"
	pb "github.com/lecex/user/proto/frontPermit"
)

//FrontPermit 仓库接口
type FrontPermit interface {
	Create(frontPermit *pb.FrontPermit) (*pb.FrontPermit, error)
	Delete(frontPermit *pb.FrontPermit) (bool, error)
	Update(frontPermit *pb.FrontPermit) (bool, error)
	Get(frontPermit *pb.FrontPermit) (*pb.FrontPermit, error)
	All(req *pb.Request) ([]*pb.FrontPermit, error)
	List(req *pb.ListQuery) ([]*pb.FrontPermit, error)
	Total(req *pb.ListQuery) (int64, error)
}

// FrontPermitRepository 前端权限仓库
type FrontPermitRepository struct {
	DB *gorm.DB
}

// All 获取所有角色信息
func (repo *FrontPermitRepository) All(req *pb.Request) (frontPermits []*pb.FrontPermit, err error) {
	if err := repo.DB.Find(&frontPermits).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return frontPermits, nil
}

// List 获取所有前端权限信息
func (repo *FrontPermitRepository) List(req *pb.ListQuery) (frontPermits []*pb.FrontPermit, err error) {
	db := repo.DB
	limit, offset := uitl.Page(req.Limit, req.Page) // 分页
	sort := uitl.Sort(req.Sort)                     // 排序 默认 created_at desc
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&frontPermits).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return frontPermits, nil
}

// Total 获取所有前端权限查询总量
func (repo *FrontPermitRepository) Total(req *pb.ListQuery) (total int64, err error) {
	frontPermits := []pb.FrontPermit{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&frontPermits).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取前端权限信息
func (repo *FrontPermitRepository) Get(frontPermit *pb.FrontPermit) (*pb.FrontPermit, error) {
	if err := repo.DB.Where(&frontPermit).Find(&frontPermit).Error; err != nil {
		return nil, err
	}
	return frontPermit, nil
}

// Create 创建前端权限
// bug 无前端权限名创建前端权限可能引起 bug
func (repo *FrontPermitRepository) Create(p *pb.FrontPermit) (*pb.FrontPermit, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return p, fmt.Errorf("添加前端权限失败")
	}
	return p, nil
}

// Update 更新前端权限
func (repo *FrontPermitRepository) Update(p *pb.FrontPermit) (bool, error) {
	if p.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.FrontPermit{
		Id: p.Id,
	}
	err := repo.DB.Model(id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除前端权限
func (repo *FrontPermitRepository) Delete(p *pb.FrontPermit) (bool, error) {
	err := repo.DB.Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

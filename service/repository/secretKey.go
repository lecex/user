package repository

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-micro/v2/util/log"

	"github.com/lecex/core/util"
	pb "github.com/lecex/user/proto/secretKey"

	"github.com/jinzhu/gorm"
)

//SecretKey 仓库接口
type SecretKey interface {
	Create(secretKey *pb.SecretKey) (*pb.SecretKey, error)
	Delete(secretKey *pb.SecretKey) (bool, error)
	Update(secretKey *pb.SecretKey) (bool, error)
	Get(secretKey *pb.SecretKey) (*pb.SecretKey, error)
	All(req *pb.Request) ([]*pb.SecretKey, error)
	List(req *pb.ListQuery) ([]*pb.SecretKey, error)
	Total(req *pb.ListQuery) (int64, error)
}

// SecretKeyRepository 秘钥仓库
type SecretKeyRepository struct {
	DB *gorm.DB
}

// All 获取所有秘钥信息
func (repo *SecretKeyRepository) All(req *pb.Request) (secretKeys []*pb.SecretKey, err error) {
	if err := repo.DB.Find(&secretKeys).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return secretKeys, nil
}

// List 获取所有秘钥信息
func (repo *SecretKeyRepository) List(req *pb.ListQuery) (secretKeys []*pb.SecretKey, err error) {
	db := repo.DB
	limit, offset := util.Page(req.Limit, req.Page) // 分页
	sort := util.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&secretKeys).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return secretKeys, nil
}

// Total 获取所有秘钥查询总量
func (repo *SecretKeyRepository) Total(req *pb.ListQuery) (total int64, err error) {
	secretKeys := []pb.SecretKey{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&secretKeys).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取秘钥信息
func (repo *SecretKeyRepository) Get(secretKey *pb.SecretKey) (*pb.SecretKey, error) {
	if err := repo.DB.Where(&secretKey).Find(&secretKey).Error; err != nil {
		return nil, err
	}
	return secretKey, nil
}

// Create 创建秘钥
// bug 无秘钥名创建秘钥可能引起 bug
func (repo *SecretKeyRepository) Create(r *pb.SecretKey) (*pb.SecretKey, error) {
	err := repo.DB.Create(r).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return r, fmt.Errorf("添加秘钥失败")
	}
	return r, nil
}

// Update 更新秘钥
func (repo *SecretKeyRepository) Update(r *pb.SecretKey) (bool, error) {
	if r.UserId == "" {
		return false, fmt.Errorf("请传入更新user_id")
	}
	err := repo.DB.Model(&r).Where("user_id = ?", r.UserId).Updates(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除秘钥
func (repo *SecretKeyRepository) Delete(k *pb.SecretKey) (bool, error) {
	if k.UserId == "" {
		return false, fmt.Errorf("请传入更新user_id")
	}
	err := repo.DB.Model(&k).Where("user_id = ?", k.UserId).Delete(k).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

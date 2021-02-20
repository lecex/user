package migrations

import (
	"context"

	"github.com/lecex/core/env"
	"github.com/micro/go-micro/v2/util/log"

	casbinPB "github.com/lecex/user/proto/casbin"
	frontPermitPB "github.com/lecex/user/proto/frontPermit"
	permissionPB "github.com/lecex/user/proto/permission"
	rolePB "github.com/lecex/user/proto/role"
	userPB "github.com/lecex/user/proto/user"

	"github.com/lecex/user/handler"
	"github.com/lecex/user/providers/casbin"
	db "github.com/lecex/user/providers/database"
	"github.com/lecex/user/service/repository"
)

func init() {
	user()
	frontPermit()
	permission()
	role()

	seeds()
}

// user 用户数据迁移
func user() {
	user := &userPB.User{}
	if !db.DB.HasTable(&user) {
		db.DB.Exec(`
			CREATE TABLE users (
			id varchar(36) NOT NULL,
			username varchar(64) DEFAULT NULL,
			mobile varchar(11) DEFAULT NULL,
			email varchar(64) DEFAULT NULL,
			password varchar(128) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			avatar varchar(128) DEFAULT NULL,
			origin varchar(32) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY username (username)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// frontPermit 前端权限数据迁移
func frontPermit() {
	frontPermit := &frontPermitPB.FrontPermit{}
	if !db.DB.HasTable(&frontPermit) {
		db.DB.Exec(`
			CREATE TABLE front_permits (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			app varchar(64) DEFAULT NULL,
			service varchar(64) DEFAULT NULL,
			method varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			description varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY service_OR_method (service,method)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// permission 权限数据迁移
func permission() {
	permission := &permissionPB.Permission{}
	if !db.DB.HasTable(&permission) {
		db.DB.Exec(`
			CREATE TABLE permissions (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			service varchar(64) DEFAULT NULL,
			method varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			description varchar(128) DEFAULT NULL,
			auth tinyint(1) DEFAULT 0,
			policy tinyint(1) DEFAULT 0,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY service_OR_method (service,method)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// role 角色数据迁移
func role() {
	role := &rolePB.Role{}
	if !db.DB.HasTable(&role) {
		db.DB.Exec(`
			CREATE TABLE roles (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			label varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			description varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY label (label)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

func seeds() {
	CreateRole()
	CreateUser()
}

func CreateRole() {
	// 角色服务实现
	repo := &repository.RoleRepository{db.DB}
	h := handler.Role{repo}
	req := &rolePB.Request{
		Role: &rolePB.Role{
			Label:       `root`,
			Name:        `超级管理员`,
			Description: `超级管理员拥有全部权限`,
		},
	}
	res := &rolePB.Response{}
	err := h.Create(context.TODO(), req, res)
	// AddRole
	log.Log(err)
}

// CreateUser 填充文件
func CreateUser() {
	password := env.Getenv("ADMIN_PASSWORD", "admin123")
	repo := &repository.UserRepository{db.DB}
	h := handler.User{repo}
	req := &userPB.Request{
		User: &userPB.User{
			Username: `admin`,
			Password: password,
			Origin:   `User`,
		},
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	if err == nil {
		// 增加用户 root 权限
		addRole(res.User.Id, `root`)
	}
	// AddRole
	log.Log(err)
}

// AddRole 增加用户角色
func addRole(userID string, role string) {
	h := handler.Casbin{casbin.Enforcer}
	req := &casbinPB.Request{
		Label: userID,
		Role:  role,
	}
	res := &casbinPB.Response{}
	err := h.AddRole(context.TODO(), req, res)
	log.Log(err)
}

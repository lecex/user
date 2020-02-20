package main

import (
	PB "github.com/lecex/user/proto/permission"
)

// Config 配置
type Config struct {
	Service     string
	Version     string
	Permissions []*PB.Permission
}

// Conf 配置
var Conf Config = Config{
	Service: "user",
	Version: "latest",
}

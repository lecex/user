package service

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/lecex/core/env"
	auth "github.com/lecex/user/proto/auth"
)

// 定义加盐哈希密码时所用的盐
// 要保证其生成和保存都足够安全
// 比如使用 md5 来生成
var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something
	privateKey = []byte(env.Getenv("APP_KEY", "8ca96774aadf77668e42931b9c0a14e5"))
)

// Authable 授权加密解密
type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *auth.User) (string, error)
}

// CustomClaims 自定义的 metadata
// 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *auth.User
	// 使用标准的 payload
	jwt.StandardClaims
}

// TokenService 令牌服务
type TokenService struct {
}

// Decode 将 JWT 字符串解密为 CustomClaims 对象
func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode 将 User 用户信息加密为 JWT 字符串
func (srv *TokenService) Encode(user *auth.User) (string, error) {
	// 默认三天后过期
	validityPeriod, _ := strconv.ParseInt(env.Getenv("TOKEN_VALIDITY_PERIOD", "3"), 10, 64)
	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(validityPeriod)).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    `user`, // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}

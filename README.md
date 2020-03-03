# user 用户

## MACOS开发环境
```
    # enable go modules
    export GO111MODULE=on

    go get github.com/micro/micro/v2

    # protobuf 开发环境
    go get -u github.com/golang/protobuf/protoc-gen-go
    go get -u github.com/micro/protoc-gen-micro/v2

    go get -u github.com/gogo/protobuf/proto
    go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
    go get -u github.com/gogo/protobuf/gogoproto
```
## user 用户服务
- Exist     查询用户
- List      用户列表
- Get       查询用户
- Create    创建用户
- Update    更新用户
- Delete    删除用户
## auth 认证服务
- Auth              用户授权
- ValidateToken     授权认证
- (**高危-内部调用**)
- AuthById          通过用户ID获取授权

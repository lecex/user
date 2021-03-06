.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
tag:
	git push --tags
.PHONY: rpc
rpc:
	micro api  --handler=rpc  --namespace=go.micro.api --address=:8080
.PHONY: api
api:
	micro api  --handler=api  --namespace=go.micro.api --address=:8081
micro:
	micro api --enable_rpc=true

.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/user/user.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/secretKey/secretKey.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/auth/auth.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/frontPermit/frontPermit.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/permission/permission.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/role/role.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/casbin/casbin.proto
	protoc -I . --micro_out=. --gogofaster_out=. proto/log/log.proto

.PHONY: docker
docker:
	docker build -f Dockerfile  -t user .
.PHONY: run
run:
	go run main.go
test:
	go test main_test.go -test.v
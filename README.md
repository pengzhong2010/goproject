# Goproject

go语言大仓开发模式仓库

# Quickstart
## help
```shell
sh tool.sh help
```
## init env 
```shell
sh tool.sh xxx init
```
xxx don't care , maybe any string
## new app server
```shell
sh tool.sh {myapp} new
```
## add new file api 
```shell
sh tool.sh {myapp} add {myapi}
```
## edit api
edit file ./api/{myapp}/v1/{myapi}.proto
## make proto file 
```shell
sh tool.sh {myapp} apiclient {myapi}
```
## make proto error file
```shell
sh tool.sh {myapp} apierror {myapierror}
```
## edit config file
### edit config file 

path: ./app/{myapp}/configs/{myapp.yaml}
### edit config proto 

path: ./app/{myapp}/internal/conf/conf.proto
### generate proto to go file
```shell
sh tool.sh {myapp} config
```
## add new file service
sh tool.sh {myapp} apiserver {myapi}
## edit service file
edit file ./app/{myapp}/internal/service/{myapi}.go
## edit server file
### http : 
file path : ./app/{myapp}/internal/server/http.go

route your server hear
```shell
v1.Register{myapp}HTTPServer(srv, {myapi})
# v1.RegisterGreeterHTTPServer(srv, greeter)
```
### grpc
file path : ./app/{myapp}/internal/server/grpc.go

route your server hear
```shell
v1.Register{myapp}Server(srv, {myapi})
#v1.RegisterGreeterServer(srv, greeter)
```
## generate wire
### edit wire Provider file
file list : ./app/{myapp}/cmd/{myapp}/wire.go
### generate wire
```shell
sh tool.sh {myapp} generate
```
## edit bussiness file 
path : ./app/{myapp}/internal/biz

## edit repo file
path : ./app/{myapp}/internal/data
## start server
```shell
ENV=local sh tool.sh {myapp} start
#ENV=dev sh tool.sh {myapp} start
```
# mysql demo
path: app/thinktank/internal/data/todo.go: Save()
# grpc request demo
path: app/thinktank/internal/biz/todo.go: Update()
# swagger ui 
path: http://{host}:{port}/q/swagger-ui
# base layout doc
- https://go-kratos.dev/docs/
- https://developers.google.com/protocol-buffers/docs/overview
- https://github.com/grpc-ecosystem/grpc-gateway
- https://pkg.go.dev/github.com/google/wire#section-readme
- https://gorm.io/docs/
- https://pkg.go.dev/github.com/gomodule/redigo/redis

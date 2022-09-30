# Goproject

go语言大仓开发模式仓库

# Quickstart
command need `shell` environment

if `windows` recommend `git-bash` for cmdline tool
## help
```shell
./ase help
```
## new app add run
```shell
./ase all {mypp}
./ase start {myapp}
```
# Useful command
## init env 
```shell
./ase init
```
## new app server
```shell
./ase new {myapp}
```
## add new file api 
```shell
./ase add {myapp} {myapi}
```
## edit api
edit file ./api/{myapp}/v1/{myapi}.proto
## make proto file 
```shell
./ase apiclient {myapp} {myapi}
```
## make proto error file
```shell
./ase apierror
```
## edit config file
### edit config file 

path: ./app/{myapp}/configs/{myapp.yaml}
### edit config proto 

path: ./app/{myapp}/internal/conf/conf.proto
### generate proto to go file
```shell
./ase config {myapp}
```
## add new file service
./ase apiserver {myapp} {myapi}
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
./ase generate {myapp}
```
## edit bussiness file 
path : ./app/{myapp}/internal/biz

## edit repo file
path : ./app/{myapp}/internal/data
## start server
```shell
# env local
ENV=local ./ase start {myapp}
# env dev
ENV=dev ./ase start {myapp}
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

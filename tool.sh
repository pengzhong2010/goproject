#!/bin/bash
set -e
#set -x

needhelp(){
  echo ''
  echo 'Usage:'
  echo ' sh tool.sh [serviceName] [command] [args]'
  echo ''
  echo 'Detail:'
  echo ' Command                Des                                eg.
 init                   init env                           sh tool.sh thinktank init
 config                 generate internal proto            sh tool.sh thinktank config
 api                    generate all api proto             sh tool.sh thinktank api
 add                    add api demo proto                 sh tool.sh thinktank add demoxxx
 apiclient              generate one api proto             sh tool.sh thinktank apiclient demoxxx
 apiserver              generate one service file          sh tool.sh thinktank apiserver demoxxx
 buildlocal             build for local system             sh tool.sh thinktank buildlocal
 buildlinux             build for linux                    sh tool.sh thinktank buildlinux
 generate               generate                           sh tool.sh thinktank generate
 all                    generate all                       sh tool.sh thinktank all
 start                  start service                      sh tool.sh thinktank start
 help                   show help                          sh tool.sh thinktank help
  '
}

if [ $# -lt 2 ]; then
  needhelp
  exit
fi

GOHOSTOS=`go env GOHOSTOS`
GOPATH=`go env GOPATH`
VERSION=`git describe --tags --always`

serviceName=$1
cmd=$2
arg=$3

INTERNAL_PROTO_FILES=`find ./app/${serviceName}/internal -name "*.proto"`
API_PROTO_FILES=`find ./api -name "*.proto"`

init(){
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go mod download
}

config(){
  protoc --proto_path=./app/${serviceName}/internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./app/${serviceName}/internal \
	       ${INTERNAL_PROTO_FILES}
  exit
}

api(){
  protoc --proto_path=. \
         --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --openapi_out=fq_schema_naming=true,default_response=false:./doc \
	       ${API_PROTO_FILES}
  exit
}

#doc(){
#  mkdir -p ./doc
#  protoc --proto_path=./api \
#	       --proto_path=./third_party \
# 	       --go_out=paths=source_relative:./api \
# 	       --go-http_out=paths=source_relative:./api \
# 	       --go-grpc_out=paths=source_relative:./api \
#	       --openapi_out=fq_schema_naming=true,default_response=false:./doc \
#	       ${API_PROTO_FILES}
#  exit
#}

buildlocal(){
  mkdir -p ./bin && go build -ldflags "-X main.Version=${VERSION}" -o ./bin/${serviceName} ./app/${serviceName}/cmd/${serviceName}/main.go ./app/${serviceName}/cmd/${serviceName}/wire_gen.go
  exit
}

buildlinux(){
  mkdir -p ./bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}" -o ./bin/${serviceName} ./app/${serviceName}/cmd/${serviceName}/main.go ./app/${serviceName}/cmd/${serviceName}/wire_gen.go
  exit
}

generate(){
  go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...
	exit
}

start(){
  cd ./app/${serviceName}/cmd/${serviceName}/
  go run main.go wire_gen.go
  exit
}

add(){
  kratos proto add api/${serviceName}/v1/${arg}.proto
  exit
}

apiclient(){
  protoc --proto_path=. \
         --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --validate_out="lang=go:../" \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --openapi_out=fq_schema_naming=true,default_response=false:./doc \
	       api/${serviceName}/v1/${arg}.proto
  exit
#  kratos proto client api/${serviceName}/v1/${arg}.proto
#  exit
}

apiserver(){
  kratos proto server api/${serviceName}/v1/${arg}.proto -t app/${serviceName}/internal/service
  exit
}


case ${cmd} in
  (init)
    init
    ;;
  (config)
    config
    ;;
  (api)
    api
    ;;
#  (doc)
#    doc
#    ;;
  (buildlocal)
    buildlocal
    ;;
  (buildlinux)
    buildlinux
    ;;
  (generate)
    generate
    ;;
  (start)
    start
    ;;
  (add)
    if [ $# -ne 3 ]; then
      echo "args {3} not found"
      exit
    fi
    add
    ;;
  (apiclient)
    if [ $# -ne 3 ]; then
      echo "args {3} not found"
      exit
    fi
    apiclient
    ;;
  (apiserver)
    if [ $# -ne 3 ]; then
      echo "args {3} not found"
      exit
    fi
    apiserver
    ;;
  (all)
    init
    api
    generate
    ;;
  (*)
    needhelp
    ;;
esac
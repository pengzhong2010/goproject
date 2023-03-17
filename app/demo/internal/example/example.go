package example

import (
	"context"
	"log"
	"time"

	pb "goproject/api/demo/v1"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"
)

func main() {
	//listTreeGrpc()
	listTreeHttp()
}
func listTreeGrpc() {
	cc, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect", err)
	}
	defer cc.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	client := pb.NewTodoClient(cc)
	resp, err := client.ListTodo(ctx, &pb.ListTodoRequest{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}

func listTreeHttp() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	cc, err := http.NewClient(ctx, http.WithEndpoint("http://127.0.0.1:8000"))
	client := pb.NewTodoHTTPClient(cc)
	resp, err := client.ListTodo(ctx, &pb.ListTodoRequest{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}

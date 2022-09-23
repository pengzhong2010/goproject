package biz

import (
	"context"
	"fmt"

	pb "goproject/api/demo/v1"
	v1 "goproject/api/helloworld/v1"
	"goproject/app/demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Todo struct {
	Id     int64
	Title  string
	Detail string
}

type TodoRepo interface {
	Save(context.Context, *Todo) error
	Update(context.Context, *Todo) error
}

// TodoUsecase is a Greeter usecase.
type TodoUsecase struct {
	ac   *conf.Auth
	repo TodoRepo
	log  *log.Helper
	tc   v1.GreeterClient
}

// NewTodoUsecase new a Greeter usecase.
func NewTodoUsecase(ac *conf.Auth, repo TodoRepo, tc v1.GreeterClient, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{ac: ac, repo: repo, tc: tc, log: log.NewHelper(logger)}
}

// CreateTodo creates a Todo, and returns the new Todo.
func (uc *TodoUsecase) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (out *pb.UpdateResp, err error) {
	err = uc.repo.Save(ctx, &Todo{
		Title:  in.Title,
		Detail: in.Detail,
	})
	if err != nil {
		err = errors.Wrap(err, "xxx")
		return
	}
	out = &pb.UpdateResp{}
	return
}
func (uc *TodoUsecase) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (out *pb.UpdateResp, err error) {
	resp, err := uc.tc.SayHello(ctx, &v1.HelloRequest{
		Name: "aax",
	})
	if err != nil {
		err = errors.Wrap(err, "request err")
		return
	}
	fmt.Println(resp)
	return
}
func (uc *TodoUsecase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	// generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 2,
	})
	signedString, err := claims.SignedString([]byte(uc.ac.Jwt.ApiSecretKey))
	if err != nil {
		err = errors.Wrap(err, "generate token failed")
		return nil, pb.ErrorAuthrizationFaild("generate token failed: %s", err.Error())
	}
	return &pb.LoginReply{
		Token: fmt.Sprintf("Bearer %s", signedString),
	}, nil
}

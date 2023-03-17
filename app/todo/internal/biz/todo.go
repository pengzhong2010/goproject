package biz

import (
	"context"
	commonpb "goproject/api/common"
	pb "goproject/api/todo/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type Todo struct {
	Id     int64
	Title  string
	Detail string
}

type TodoRepo interface {
	Create(context.Context, *pb.CreateTodoRequest) (*pb.CreateTodoReply, error)
	Update(context.Context, *pb.UpdateTodoRequest) error
	Get(context.Context, int64) (*pb.GetTodoReply, error)
	Delete(context.Context, int64) error
	List(context.Context) (*pb.ListTodoReply, error)
}

// TodoUsecase is a Greeter usecase.
type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

// NewTodoUsecase new a Greeter usecase.
func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "biz/todo"))}
}

// CreateTodo creates a Todo, and returns the new Todo.
func (uc *TodoUsecase) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (out *pb.CreateTodoReply, err error) {
	out, err = uc.repo.Create(ctx, in)
	if err != nil {
		return
	}
	return
}
func (uc *TodoUsecase) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (out *commonpb.UpdateResp, err error) {
	err = uc.repo.Update(ctx, in)
	if err != nil {
		return
	}
	out = &commonpb.UpdateResp{}
	return
}
func (uc *TodoUsecase) DeleteTodo(ctx context.Context, in int64) (out *commonpb.UpdateResp, err error) {
	err = uc.repo.Delete(ctx, in)
	if err != nil {
		return
	}
	out = &commonpb.UpdateResp{}
	return
}
func (uc *TodoUsecase) GetTodo(ctx context.Context, in int64) (out *pb.GetTodoReply, err error) {
	out, err = uc.repo.Get(ctx, in)
	return
}
func (uc *TodoUsecase) ListTodo(ctx context.Context) (out *pb.ListTodoReply, err error) {
	out, err = uc.repo.List(ctx)
	return
}

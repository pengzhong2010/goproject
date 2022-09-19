package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Todo struct {
	Id     int64
	Title  string
	Detail string
}

// GreeterRepo is a Greater repo.
type TodoRepo interface {
	Save(context.Context, *Todo) (*Todo, error)
	Update(context.Context, *Todo) (*Todo, error)
	FindByID(context.Context, int64) (*Todo, error)
	ListByHello(context.Context, string) ([]*Todo, error)
	ListAll(context.Context) ([]*Todo, error)
}

// GreeterUsecase is a Greeter usecase.
type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *TodoUsecase) CreateTodo(ctx context.Context, g *Todo) (*Todo, error) {
	uc.log.WithContext(ctx).Infof("CreateTodo: %v", g)
	return uc.repo.Save(ctx, g)
}

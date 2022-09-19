package data

import (
	"context"

	"goproject/app/thinktank/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type todoRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todoRepo) Save(ctx context.Context, in *biz.Todo) (out *biz.Todo, err error) {
	tx := r.data.Db.Begin()
	defer func() {
		if err != nil && tx != nil {
			tx.Rollback()
		}
	}()
	// 创建
	t := &Todolist{
		Title:  in.Title,
		Detail: in.Detail,
	}
	err = tx.Create(t).Error
	if err != nil {
		err = errors.Wrap(err, "Create Model err")
		return
	}
	// commit
	if err = tx.Commit().Error; err != nil {
		err = errors.Wrap(err, "commit err")
		return
	}
	return
}

func (r *todoRepo) Update(ctx context.Context, g *biz.Todo) (*biz.Todo, error) {
	return g, nil
}

func (r *todoRepo) FindByID(context.Context, int64) (*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) ListByHello(context.Context, string) ([]*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) ListAll(context.Context) ([]*biz.Todo, error) {
	return nil, nil
}

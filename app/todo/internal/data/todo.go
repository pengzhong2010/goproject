package data

import (
	"context"
	"time"

	pb "goproject/api/todo/v1"
	"goproject/app/todo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todoRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/todo")),
	}
}

func (r *todoRepo) Create(ctx context.Context, in *pb.CreateTodoRequest) (out *pb.CreateTodoReply, err error) {
	t := &TodoTable{
		Title:  in.Title,
		Detail: in.Detail,
	}
	err = r.data.db.Create(t).Error
	if err != nil {
		return
	}
	out = &pb.CreateTodoReply{Id: int32(t.ID)}
	return
}

func (r *todoRepo) Update(ctx context.Context, in *pb.UpdateTodoRequest) (err error) {
	todo := &TodoTable{
		ID:     int64(in.Id),
		Title:  in.Title,
		Detail: in.Detail,
	}
	return r.data.db.Save(todo).Error
}

func (r *todoRepo) Get(ctx context.Context, in int64) (out *pb.GetTodoReply, err error) {
	todo := &TodoTable{}
	if err = r.data.db.Where("id=?", in).Where("is_deleted=0").Select("id", "title", "detail").First(todo).Error; err != nil {
		return
	}
	out = &pb.GetTodoReply{
		Id:     int32(todo.ID),
		Title:  todo.Title,
		Detail: todo.Detail,
	}
	return
}
func (r *todoRepo) Delete(ctx context.Context, in int64) (err error) {
	todo := &TodoTable{}
	if err = r.data.db.Where("id=?", in).Where("is_deleted=0").First(todo).Error; err != nil {
		return
	}
	todo.IsDeleted = time.Now().Unix()
	return r.data.db.Save(todo).Error
}

func (r *todoRepo) List(context.Context) (out *pb.ListTodoReply, err error) {
	var ts []*TodoTable
	if err = r.data.db.Where("is_deleted=0").Find(&ts).Error; err != nil {
		return
	}
	out = &pb.ListTodoReply{}
	for _, i := range ts {
		out.Data = append(out.Data, &pb.TodoForList{
			Id:    int32(i.ID),
			Title: i.Title,
		})
	}
	return
}

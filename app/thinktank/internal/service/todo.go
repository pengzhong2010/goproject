package service

import (
	"context"
	"goproject/app/thinktank/internal/biz"

	pb "goproject/api/thinktank/v1"
)

type TodoService struct {
	pb.UnimplementedTodoServer

	uc *biz.TodoUsecase
}

func NewTodoService(uc *biz.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (resp *pb.UpdateResp, err error) {
	err = req.Validate()
	if err != nil {
		return
	}
	_, err = s.uc.CreateTodo(ctx, &biz.Todo{
		Title:  req.Title,
		Detail: req.Detail,
	})
	if err != nil {
		return
	}
	resp = &pb.UpdateResp{}
	return
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.IDReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *TodoService) GetTodo(ctx context.Context, req *pb.IDReq) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}

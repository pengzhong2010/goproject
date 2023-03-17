package service

import (
	"context"
	pbcommon "goproject/api/common"
	pb "goproject/api/todo/v1"
	"goproject/app/todo/internal/biz"
)

type TodoService struct {
	pb.UnimplementedTodoServer

	b *biz.TodoUsecase
}

func NewTodoService(b *biz.TodoUsecase) *TodoService {
	return &TodoService{b: b}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	//return &pb.CreateTodoReply{
	//	Id: 1,
	//}, nil
	return s.b.CreateTodo(ctx, req)
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pbcommon.UpdateResp, error) {
	return s.b.UpdateTodo(ctx, req)
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pbcommon.IDReq) (*pbcommon.UpdateResp, error) {
	return s.b.DeleteTodo(ctx, req.Id)
}
func (s *TodoService) GetTodo(ctx context.Context, req *pbcommon.IDReq) (*pb.GetTodoReply, error) {
	return s.b.GetTodo(ctx, req.Id)
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return s.b.ListTodo(ctx)
}

package service

import (
	"context"

	commonpb "goproject/api/common"
	pb "goproject/api/demo/v1"
	"goproject/app/demo/internal/biz"
)

type TodoService struct {
	pb.UnimplementedTodoServer

	uc *biz.TodoUsecase
}

func NewTodoService(uc *biz.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (resp *commonpb.UpdateResp, err error) {
	return s.uc.CreateTodo(ctx, req)
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (resp *commonpb.UpdateResp, err error) {
	return s.uc.UpdateTodo(ctx, req)
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *commonpb.IDReq) (*commonpb.UpdateResp, error) {
	return &commonpb.UpdateResp{}, nil
}
func (s *TodoService) GetTodo(ctx context.Context, req *commonpb.IDReq) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}
func (s *TodoService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return s.uc.Login(ctx, req)
}

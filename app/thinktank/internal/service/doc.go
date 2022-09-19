package service

import (
	"context"
	"fmt"

	pb "goproject/api/thinktank/v1"
)

type DocService struct {
	pb.UnimplementedDocServer
}

func NewDocService() *DocService {
	return &DocService{}
}

func (s *DocService) CreateDoc(ctx context.Context, req *pb.CreateDocRequest) (*pb.UpdateResp, error) {
	fmt.Println("123ok")
	fmt.Printf("%+v", req)
	return &pb.UpdateResp{}, nil
}
func (s *DocService) UpdateDoc(ctx context.Context, req *pb.UpdateDocRequest) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *DocService) DeleteDoc(ctx context.Context, req *pb.IDReq) (*pb.UpdateResp, error) {
	return &pb.UpdateResp{}, nil
}
func (s *DocService) GetDoc(ctx context.Context, req *pb.IDReq) (*pb.GetDocReply, error) {
	return &pb.GetDocReply{}, nil
}
func (s *DocService) ListDoc(ctx context.Context, req *pb.ListDocRequest) (*pb.ListDocReply, error) {
	return &pb.ListDocReply{}, nil
}

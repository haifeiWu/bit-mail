package service

import (
	"context"
	"encoding/json"
	
	pb "bit-mail/api/user/v1"
	
	v1 "bit-mail/api/user/v1"
	"bit-mail/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	data, err := s.uc.ListUser(ctx)
	if err != nil {
		return nil, err
	}
	strByte, _ := json.Marshal(data)
	return &pb.ListUserReply{
		Message: string(strByte),
	}, nil
}

func (s *UserService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{
		Stat:    0,
		Code:    0,
		Message: "success",
		Data:    "success",
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return s.uc.Login(ctx, req.GetUsername(), req.GetPassword())
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return s.uc.Register(ctx, req)
}
func (s *UserService) GetContactListByUserId(ctx context.Context, req *pb.GetContactListByUserIdRequest) (*pb.GetContactListByUserIdReply, error) {
	return s.uc.GetContactListByUserId(ctx, req)
}
func (s *UserService) UploadContact(ctx context.Context, req *pb.UploadContactRequest) (*pb.UploadContactReply, error) {
	return s.uc.UploadContact(ctx, req)
}
func (s *UserService) DelContactByUserID(ctx context.Context, req *pb.DelContactByUserIDRequest) (*pb.DelContactByUserIDReply, error) {
	return s.uc.DelContactByUserID(ctx, req)
}

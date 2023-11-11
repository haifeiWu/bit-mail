package service

import (
	"context"
	
	pb "bit-mail/api/mail/v1"
	"bit-mail/internal/biz"
)

type MailMessageService struct {
	pb.UnimplementedMailMessageServer
	uc *biz.UserUsecase
}

func NewMailMessageService(uc *biz.UserUsecase) *MailMessageService {
	return &MailMessageService{uc: uc}
}

func (s *MailMessageService) ListMailMessageByUserID(ctx context.Context, req *pb.ListMailMessageByUserIDRequest) (*pb.ListMailMessageByUserIDReply, error) {
	return &pb.ListMailMessageByUserIDReply{}, nil
}

func (s *MailMessageService) AddMailMessageByUserID(ctx context.Context, req *pb.AddMailMessageByUserIDRequest) (*pb.AddMailMessageByUserIDReply, error) {
	return &pb.AddMailMessageByUserIDReply{}, nil
}

func (s *MailMessageService) UpdateMailMessageByUserID(ctx context.Context, req *pb.UpdateMailMessageByUserIDRequest) (*pb.UpdateMailMessageByUserIDReply, error) {
	return &pb.UpdateMailMessageByUserIDReply{}, nil
}

func (s *MailMessageService) DelMailMessageByUserID(ctx context.Context, req *pb.DelMailMessageByUserIDRequest) (*pb.DelMailMessageByUserIDReply, error) {
	return &pb.DelMailMessageByUserIDReply{}, nil
}

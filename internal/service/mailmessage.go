package service

import (
	"context"
	
	pb "bit-mail/api/mail/v1"
	"bit-mail/internal/biz"
)

type MailMessageService struct {
	pb.UnimplementedMailMessageServer
	uc *biz.MailUsecase
}

func NewMailMessageService(uc *biz.MailUsecase) *MailMessageService {
	return &MailMessageService{uc: uc}
}

func (s *MailMessageService) ListMailMessageByUserID(ctx context.Context, req *pb.ListMailMessageByUserIDRequest) (*pb.ListMailMessageByUserIDReply, error) {
	return s.uc.ListMailMessageByUserID(ctx, req)
}

func (s *MailMessageService) AddMailMessageByUserID(ctx context.Context, req *pb.AddMailMessageByUserIDRequest) (*pb.AddMailMessageByUserIDReply, error) {
	return s.uc.AddMailMessageByUserID(ctx, req)
}

func (s *MailMessageService) UpdateMailMessageByUserID(ctx context.Context, req *pb.UpdateMailMessageByUserIDRequest) (*pb.UpdateMailMessageByUserIDReply, error) {
	return s.uc.UpdateMailMessageByUserID(ctx, req)
}

func (s *MailMessageService) DelMailMessageByUserID(ctx context.Context, req *pb.DelMailMessageByUserIDRequest) (*pb.DelMailMessageByUserIDReply, error) {
	return s.uc.DelMailMessageByUserID(ctx, req)
}

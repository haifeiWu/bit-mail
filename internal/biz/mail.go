package biz

import (
	"context"
	"time"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	
	"bit-mail/api/mail/v1"
	"bit-mail/pkg/model"
)

// MailRepo is a Greater repo.
type MailRepo interface {
	AddMailMessageByUserID(ctx context.Context, userId string, dbEmail *model.Email) error
}

// MailUsecase is a Greeter usecase.
type MailUsecase struct {
	repo MailRepo
	log  *log.Helper
}

// NewMailUsecase new a Greeter usecase.
func NewMailUsecase(repo MailRepo, logger log.Logger) *MailUsecase {
	return &MailUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u MailUsecase) AddMailMessageByUserID(ctx context.Context, req *v1.AddMailMessageByUserIDRequest) (*v1.AddMailMessageByUserIDReply, error) {
	var err error
	dbEmail := &model.Email{
		Subject:   req.GetSubject(),
		Body:      req.GetBody(),
		SenderID:  cast.ToInt32(req.GetSenderId()),
		SentAt:    time.Time{},
		CcList:    req.GetCcList(),
		BccList:   req.GetBccList(),
		IsDraft:   false,
		IsDeleted: false,
	}
	
	reply := &v1.AddMailMessageByUserIDReply{
		Stat:    0,
		Code:    0,
		Message: "",
		Data:    "",
	}
	
	err = u.repo.AddMailMessageByUserID(ctx, req.GetUserId(), dbEmail)
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return nil, err
	}
	reply.Message = "success"
	return reply, nil
}

func (u MailUsecase) ListMailMessageByUserID(ctx context.Context, req *v1.ListMailMessageByUserIDRequest) (*v1.ListMailMessageByUserIDReply, error) {
	return nil, nil
}

func (u MailUsecase) UpdateMailMessageByUserID(ctx context.Context, req *v1.UpdateMailMessageByUserIDRequest) (*v1.UpdateMailMessageByUserIDReply, error) {
	return nil, nil
}

func (u MailUsecase) DelMailMessageByUserID(ctx context.Context, req *v1.DelMailMessageByUserIDRequest) (*v1.DelMailMessageByUserIDReply, error) {
	return nil, nil
}

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
	ListMailMessageByUserID(ctx context.Context, userId string, folder string, isDraft string, isDelete string) ([]model.Email, error)
	UpdateMailMessageByUserID(ctx context.Context, dbModel model.Email) error
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
	var (
		err       error
		emailList = make([]model.Email, 0)
		reply     = &v1.ListMailMessageByUserIDReply{
			Stat:    0,
			Code:    0,
			Message: "",
			Data:    make([]*v1.ListMailMessageByUserIDReply_MailMessage, 0),
		}
	)
	
	emailList, err = u.repo.ListMailMessageByUserID(ctx, req.GetUserId(), req.GetFolder(), req.GetIsDraft(), req.GetIsDelete())
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, nil
	}
	for _, e := range emailList {
		reply.Data = append(reply.Data, &v1.ListMailMessageByUserIDReply_MailMessage{
			Id:       cast.ToUint32(e.ID),
			Subject:  e.Subject,
			Body:     e.Body,
			SenderId: cast.ToUint32(e.SenderID),
			SentAt:   e.SentAt.String(),
			CcList:   e.CcList,
			BccList:  e.BccList,
		})
	}
	reply.Message = "success"
	return reply, nil
}

func (u MailUsecase) UpdateMailMessageByUserID(ctx context.Context, req *v1.UpdateMailMessageByUserIDRequest) (*v1.UpdateMailMessageByUserIDReply, error) {
	var (
		err          error
		emailDbModel model.Email
		reply        = &v1.UpdateMailMessageByUserIDReply{
			Stat:    0,
			Code:    0,
			Message: "",
			Data:    "",
		}
	)
	
	emailDbModel = model.Email{
		ID:       cast.ToInt32(req.GetMessageId()),
		Subject:  req.GetSubject(),
		Body:     req.GetBody(),
		SenderID: cast.ToInt32(req.GetSenderId()),
	}
	
	err = u.repo.UpdateMailMessageByUserID(ctx, emailDbModel)
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, nil
	}
	
	reply.Message = "success"
	return reply, nil
}

func (u MailUsecase) DelMailMessageByUserID(ctx context.Context, req *v1.DelMailMessageByUserIDRequest) (*v1.DelMailMessageByUserIDReply, error) {
	
	return nil, nil
}

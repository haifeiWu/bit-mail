package data

import (
	"context"
	
	"github.com/go-kratos/kratos/v2/log"
	
	"bit-mail/internal/biz"
	"bit-mail/pkg/model"
)

var _ biz.MailRepo = (*MailRepo)(nil)

type MailRepo struct {
	data *Data
	log  *log.Helper
}

func (m MailRepo) AddMailMessageByUserID(ctx context.Context, userId string, dbEmail *model.Email) error {
	// TODO implement me
	panic("implement me")
}

// NewUserRepo .
func NewMailRepo(data *Data, logger log.Logger) biz.MailRepo {
	return &MailRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

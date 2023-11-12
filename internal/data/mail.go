package data

import (
	"context"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	
	"bit-mail/internal/biz"
	"bit-mail/pkg/model"
)

var _ biz.MailRepo = (*MailRepo)(nil)

type MailRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewMailRepo(data *Data, logger log.Logger) biz.MailRepo {
	return &MailRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m MailRepo) ListMailMessageByUserID(ctx context.Context, userId string, folder string, isDraft string, isDelete string) ([]model.Email, error) {
	userEmail := model.UserEmail{}
	emIds := make([]string, 0)
	err := m.data.DB.WithContext(ctx).
		Table(userEmail.TableName()).
		Select("email_id,is_read").
		Joins("left join users on users.id = user_emails.recipient_id").
		Find(&emIds).Error
	if err != nil {
		return nil, err
	}
	
	email := model.Email{}
	emails := make([]model.Email, 0)
	err = m.data.DB.WithContext(ctx).Table(email.TableName()).Where("id in (?)", emIds).Find(&emails).Error
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (m MailRepo) UpdateMailMessageByUserID(ctx context.Context, dbModel model.Email) error {
	err := m.data.DB.WithContext(ctx).Table(dbModel.TableName()).
		Where("id = ?", dbModel.ID).
		Update("body = ?", dbModel.Body).
		Update("subject=?", dbModel.Subject).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MailRepo) AddMailMessageByUserID(ctx context.Context, userId string, dbEmail *model.Email) error {
	tx := m.data.DB.WithContext(ctx).Begin()
	err := tx.Table(dbEmail.TableName()).Create(&dbEmail).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	// 1：收件箱，2：发件箱，3：草稿，4：垃圾箱
	userEmail := model.UserEmail{
		EmailID:     dbEmail.ID,
		RecipientID: cast.ToInt32(userId),
		IsRead:      false,
		Folder:      2,
	}
	err = tx.Table(userEmail.TableName()).Create(&userEmail).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

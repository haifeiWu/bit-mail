package data

import (
	"context"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	
	"bit-mail/internal/biz"
	"bit-mail/pkg/model"
	"bit-mail/pkg/model_dto"
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

func (m MailRepo) ListMailMessageByUserID(ctx context.Context, userId string, folder string, isDraft string, isDelete string) ([]model_dto.UserAndEmail, error) {
	user := model.User{}
	uae := make([]model_dto.UserAndEmail, 0)
	err := m.data.DB.WithContext(ctx).
		Table(user.TableName()).
		Select("user_emails.*, emails.*").
		Where("users.id = ? and user_emails.folder = ?", userId, folder).
		Joins("INNER JOIN user_emails ON users.id = user_emails.recipient_id").
		Joins("INNER JOIN emails ON user_emails.email_id = emails.id").
		// Joins("left join users on users.id = user_emails.recipient_id").
		Find(&uae).Error
	if err != nil {
		return nil, err
	}
	
	return uae, nil
}

func (m MailRepo) UpdateMailMessageByUserID(ctx context.Context, dbModel model.Email, read bool) error {
	updateMap := make(map[string]interface{})
	if dbModel.Subject != "" {
		updateMap["subject"] = dbModel.Subject
	}
	if dbModel.Body != "" {
		updateMap["body"] = dbModel.Body
	}
	
	if dbModel.Img != "" {
		updateMap["img"] = dbModel.Img
	}
	
	if dbModel.ReceiverID != 0 {
		updateMap["receiver_id"] = dbModel.ReceiverID
	}
	
	tx := m.data.DB.WithContext(ctx).Begin()
	err := tx.Table(dbModel.TableName()).
		Where("id = ?", dbModel.ID).
		Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	if read {
		var userEmail model.UserEmail
		err = tx.Table(userEmail.TableName()).
			Where("email_id = ?", dbModel.ID).
			Update("is_read", true).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	tx.Commit()
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
		RecipientID: cast.ToInt32(dbEmail.SenderID),
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

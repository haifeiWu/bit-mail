package data

import (
	"context"
	
	"github.com/go-kratos/kratos/v2/log"
	
	"bit-mail/internal/biz"
	"bit-mail/pkg/model"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u UserRepo) ListAll(ctx context.Context) ([]*model.User, error) {
	var (
		user        model.User
		resUserList []*model.User
	)
	err := u.data.DB.WithContext(ctx).Table(user.TableName()).Select("*").Find(&resUserList).Error
	if err != nil {
		return nil, err
	}
	return resUserList, nil
}

func (u UserRepo) FindUserByName(ctx context.Context, username string) (*model.User, error) {
	var (
		user model.User
		err  error
	)
	
	err = u.data.DB.WithContext(ctx).
		Table(user.TableName()).
		Where("username = ?", username).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (u UserRepo) CreateUser(ctx context.Context, user model.User) error {
	var (
		err error
	)
	err = u.data.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) GetContactListByUserId(ctx context.Context, userId string, friendName string, friendPhone string, friendEmail string) ([]model.Contact, error) {
	var (
		err            error
		contact        = model.Contact{}
		resContactList = make([]model.Contact, 0)
	)
	db := u.data.DB.WithContext(ctx).Table(contact.TableName()).Where("user_id = ?", userId)
	if friendName != "" {
		db.Where("contact_name = ?", friendName)
	}
	
	if friendPhone != "" {
		db.Where("phone_number = ?", friendPhone)
	}
	
	if friendEmail != "" {
		db.Where("contact_email = ?", friendEmail)
	}
	
	err = db.Find(&resContactList).Error
	if err != nil {
		return nil, err
	}
	return resContactList, nil
}

func (u UserRepo) CreateContactInBatch(ctx context.Context, contacts []model.Contact) error {
	contant := model.Contact{}
	err := u.data.DB.WithContext(ctx).Table(contant.TableName()).CreateInBatches(&contacts, 10).Error
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) DeleteContactByUserID(ctx context.Context, userId string, friendIds []string) error {
	contant := model.Contact{}
	resContactList := make([]model.Contact, 0)
	tx := u.data.DB.WithContext(ctx).Table(contant.TableName()).Begin()
	
	err := tx.Where("user_id = ?", userId).
		Where("contact_user_id in (?) and is_delete = 1", friendIds).
		Find(&resContactList).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	updateFids := make([]int32, 0)
	for _, c := range resContactList {
		updateFids = append(updateFids, c.ContactUserID)
	}
	err = tx.Where("user_id = ?", userId).
		Where("contact_user_id in (?)", updateFids).
		Update("is_delete", 2).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	tx.Commit()
	return nil
}

func (u UserRepo) FindUserByID(ctx context.Context, userId string) (*model.User, error) {
	var (
		user model.User
		err  error
	)
	
	err = u.data.DB.WithContext(ctx).
		Table(user.TableName()).
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

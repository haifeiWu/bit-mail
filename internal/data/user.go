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
		First(&user).Error
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
	// var (
	// 	err error
	// )
	// err = u.data.DB.WithContext(ctx).Table()
	// TODO implement me
	panic("implement me")
}

func (u UserRepo) CreateContactInBatch(ctx context.Context, contacts []model.Contact) error {
	// TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteContactByUserID(ctx context.Context, userId string, friendIds []string) error {
	// TODO implement me
	panic("implement me")
}

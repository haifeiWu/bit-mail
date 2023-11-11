package biz

import (
	"context"
	"time"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	
	"bit-mail/api/user/v1"
	"bit-mail/pkg/model"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	ListAll(context.Context) ([]*model.User, error)
	FindUserByName(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	GetContactListByUserId(ctx context.Context, userId string, friendName string, friendPhone string, friendEmail string) ([]model.Contact, error)
	CreateContactInBatch(ctx context.Context, contacts []model.Contact) error
	DeleteContactByUserID(ctx context.Context, userId string, friendIds []string) error
	FindUserByID(ctx context.Context, userId string) (*model.User, error)
}

// UserUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) ListUser(ctx context.Context) ([]*model.User, error) {
	return uc.repo.ListAll(ctx)
}

func (uc *UserUsecase) Login(ctx context.Context, username string, password string) (*v1.LoginReply, error) {
	var (
		user  = &model.User{}
		err   error
		reply = &v1.LoginReply{
			Stat:    0,
			Code:    0,
			Message: "",
			Data:    "",
		}
	)
	
	user, err = uc.repo.FindUserByName(ctx, username)
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		uc.log.WithContext(ctx).Errorf("FindUserByNameErr=%v", err)
		return reply, nil
	}
	
	if user.Password == password {
		reply.Message = "success"
		reply.Data = cast.ToString(user.ID)
		return reply, nil
	}
	
	return reply, nil
}

func (uc *UserUsecase) Register(ctx context.Context, registerRequest *v1.RegisterRequest) (*v1.RegisterReply, error) {
	now := time.Now()
	user := model.User{
		ID:          0,
		Username:    registerRequest.GetUsername(),
		Password:    registerRequest.GetPassword(),
		Email:       registerRequest.GetEmail(),
		FirstName:   registerRequest.GetFirstName(),
		LastName:    registerRequest.GetLastName(),
		PhoneNumber: registerRequest.GetPhoneNumber(),
		CreatedAt:   now,
		UpdatedAt:   now,
		Gender:      registerRequest.GetGender(),
	}
	reply := &v1.RegisterReply{
		Stat:    0,
		Code:    0,
		Message: "",
		Data:    "",
	}
	
	alreadyExistUser, err := uc.repo.FindUserByName(ctx, user.Username)
	if alreadyExistUser != nil {
		reply.Stat = 0
		reply.Code = 2
		reply.Message = user.Username + "该用户已存在"
		reply.Data = cast.ToString(alreadyExistUser.ID)
		return reply, nil
	}
	
	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, err
	}
	
	reply.Message = "success"
	return reply, nil
}

func (uc *UserUsecase) GetContactListByUserId(ctx context.Context, req *v1.GetContactListByUserIdRequest) (*v1.GetContactListByUserIdReply, error) {
	contactList := make([]model.Contact, 0)
	var err error
	reply := &v1.GetContactListByUserIdReply{
		Stat:    0,
		Code:    0,
		Message: "",
	}
	
	contactList, err = uc.repo.GetContactListByUserId(ctx, req.GetUserId(), req.GetFriendName(), req.GetFriendPhone(), req.GetFriendEmail())
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, err
	}
	
	reply.Message = "success"
	for _, contact := range contactList {
		reapContact := &v1.GetContactListByUserIdReply_Contact{
			ContactUserId: contact.ContactUserID,
			ContactName:   contact.ContactName,
			ContactEmail:  contact.ContactEmail,
			Address:       contact.Address,
			PhoneNumber:   contact.PhoneNumber,
			Note:          contact.Note,
		}
		
		reply.Data = append(reply.Data, reapContact)
	}
	return reply, nil
}

func (uc *UserUsecase) UploadContact(ctx context.Context, uploadContactRequest *v1.UploadContactRequest) (*v1.UploadContactReply, error) {
	contactList := make([]model.Contact, 0)
	var err error
	reply := &v1.UploadContactReply{
		Stat:    0,
		Code:    0,
		Message: "",
		Data:    "",
	}
	
	userId := uploadContactRequest.GetUserId()
	reqContactList := uploadContactRequest.GetContactList()
	for _, contant := range reqContactList {
		contactList = append(contactList, model.Contact{
			UserID:       cast.ToInt32(userId),
			ContactName:  contant.GetContactName(),
			ContactEmail: contant.GetContactEmail(),
			Address:      contant.GetAddress(),
			PhoneNumber:  contant.GetPhoneNumber(),
			Note:         contant.GetNote(),
			CreatedAt:    time.Time{},
			UpdateAt:     time.Time{},
		})
	}
	
	err = uc.repo.CreateContactInBatch(ctx, contactList)
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, err
	}
	
	reply.Message = "success"
	return reply, nil
}

func (uc *UserUsecase) DelContactByUserID(ctx context.Context, delContactByUserIDRequest *v1.DelContactByUserIDRequest) (*v1.DelContactByUserIDReply, error) {
	var err error
	reply := &v1.DelContactByUserIDReply{
		Stat:    0,
		Code:    0,
		Message: "",
		Data:    "",
	}
	
	err = uc.repo.DeleteContactByUserID(ctx, delContactByUserIDRequest.GetUserId(), delContactByUserIDRequest.GetFriendIds())
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, err
	}
	
	reply.Message = "success"
	return reply, nil
}

func (uc *UserUsecase) GetUserDetailsByID(ctx context.Context, detailsByIDRequest *v1.GetUserDetailsByIDRequest) (*v1.GetUserDetailsByIDReply, error) {
	var err error
	reply := &v1.GetUserDetailsByIDReply{
		Stat:    0,
		Code:    0,
		Message: "",
	}
	user := &model.User{}
	user, err = uc.repo.FindUserByID(ctx, detailsByIDRequest.GetUserId())
	if err != nil {
		reply.Stat = -1
		reply.Code = -1
		reply.Message = "failed"
		return reply, err
	}
	
	reply.Message = "success"
	reply.Data = &v1.GetUserDetailsByIDReply_User{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		// BirthDate:    user.BirthDate,
		PhoneNumber: user.PhoneNumber,
		// ProfilePhoto: user.ProfilePhoto,
		Gender: user.Gender,
	}
	return reply, nil
}

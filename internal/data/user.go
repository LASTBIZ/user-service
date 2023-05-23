package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
	"user-service/internal/biz"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
)

type Messenger struct {
	KakaoTalk string
	Telegram  string
	WhatsApp  string
	Line      string
	Signal    string
}

type User struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsVerify  bool `gorm:"default:false"`
	Firstname string
	Lastname  string
	//FullName     string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	Blocked      bool `gorm:"default:false"`
	Role         string
	Email        string
	Phone        string
	Organization string
	Messengers   Messenger `gorm:"embedded"`
	//Messengers   []Messenger `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize int) ([]*biz.User, int, error) {
	var users []User
	result := r.data.db.Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, 0, errors.New(500, "FIND_USER_ERROR", "find user error")
	}
	total := int(result.RowsAffected)
	r.data.db.Scopes(paginate(pageNum, pageSize)).Find(&users)
	rv := make([]*biz.User, 0)
	for _, u := range users {
		rv = append(rv, modelToResponse(u))
	}
	return rv, total, nil
}

func (r *userRepo) UserByEmail(ctx context.Context, email string) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	re := modelToResponse(user)
	return re, nil
}

func (r *userRepo) GetUserById(ctx context.Context, id uint32) (*biz.User, error) {
	var user User
	if err := r.data.db.Where(&User{ID: id}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	re := modelToResponse(user)
	return re, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (bool, error) {
	var userInfo User
	result := r.data.db.Where(&User{ID: user.ID}).First(&userInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("USER_NOT_FOUND", "user not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("USER_NOT_FOUND", "rows null")
	}

	userInfo.Firstname = user.Firstname
	userInfo.Lastname = user.Lastname
	userInfo.Role = user.Role
	userInfo.IsVerify = user.IsVerify
	userInfo.Phone = user.Phone
	userInfo.Organization = user.Organization
	userInfo.Messengers = Messenger{
		KakaoTalk: user.Messengers.KakaoTalk,
		Telegram:  user.Messengers.Telegram,
		Line:      user.Messengers.Line,
		WhatsApp:  user.Messengers.WhatsApp,
		Signal:    user.Messengers.Signal,
	}
	userInfo.Blocked = user.Blocked
	userInfo.UpdatedAt = time.Now()

	if err := r.data.db.Save(&userInfo).Error; err != nil {
		return false, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return true, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint32) (bool, error) {
	var userInfo User
	result := r.data.db.Where(&User{ID: id}).First(&userInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("USER_NOT_FOUND", "user not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("USER_NOT_FOUND", "rows null")
	}

	err := r.data.db.Delete(&userInfo)
	if err.Error != nil {
		return false, errors.NotFound("USER_DELETE_ERROR", "user delete error")
	}

	return true, nil
}

func modelToResponse(user User) *biz.User {
	userInfoRsp := &biz.User{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		IsVerify:     user.IsVerify,
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Blocked:      user.Blocked,
		Role:         user.Role,
		Email:        user.Email,
		Phone:        user.Phone,
		Organization: user.Organization,
		Messengers: biz.Messenger{
			KakaoTalk: user.Messengers.KakaoTalk,
			Telegram:  user.Messengers.Telegram,
			WhatsApp:  user.Messengers.WhatsApp,
			Line:      user.Messengers.Line,
			Signal:    user.Messengers.Signal,
		},
	}
	return userInfoRsp
}

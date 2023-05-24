package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Messenger struct {
	KakaoTalk string
	Telegram  string
	WhatsApp  string
	Line      string
	Signal    string
}

type User struct {
	ID           uint32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsVerify     bool
	Firstname    string
	Lastname     string
	Blocked      bool
	Role         string
	Email        string
	Phone        string
	Organization string
	Messengers   Messenger
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	UserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id uint32) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	DeleteUser(ctx context.Context, id uint32) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) UserByEmail(ctx context.Context, email string) (*User, error) {
	return uc.repo.UserByEmail(ctx, email)
}

func (uc *UserUseCase) UserById(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint32) (bool, error) {
	return uc.repo.DeleteUser(ctx, id)
}

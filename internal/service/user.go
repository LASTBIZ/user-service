package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user-service/internal/biz"

	pb "user-service/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfoResponse, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Email:     req.Email,
		Firstname: req.FirstName,
		Lastname:  req.LastName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := UserResponse(user)
	return userInfoRsp, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*empty.Empty, error) {
	user, err := u.uc.Update(ctx, &biz.User{
		ID:           uint32(req.Id),
		Firstname:    req.FirstName,
		Lastname:     req.LastName,
		Role:         req.Role,
		IsVerify:     req.IsVerify,
		Phone:        req.Phone,
		Organization: req.Organization,
		Messengers: biz.Messenger{
			KakaoTalk: req.Messenger.KakaoTalk,
			Telegram:  req.Messenger.Telegram,
			Line:      req.Messenger.Line,
			WhatsApp:  req.Messenger.WhatsApp,
			Signal:    req.Messenger.Signal,
		},
		Blocked: req.Blocked,
	})

	if err != nil {
		return nil, err
	}

	if user == false {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.IdRequest) (*empty.Empty, error) {
	user, err := u.uc.DeleteUser(ctx, uint32(req.Id))
	if err != nil {
		return nil, err
	}

	if user == false {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserInfoResponse, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user id")
	defer span.End()
	user, err := u.uc.UserById(ctx, uint32(req.Id))
	if err != nil {
		return nil, err
	}
	rsp := UserResponse(user)
	return rsp, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, req *pb.EmailRequest) (*pb.UserInfoResponse, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user email")
	defer span.End()
	user, err := u.uc.UserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	rsp := UserResponse(user)
	return rsp, nil
}

func (u *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.UserListResponse, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user list")
	defer span.End()

	list, total, err := u.uc.List(ctx, int(req.Pn), int(req.PSize))
	if err != nil {
		return nil, err
	}
	rsp := &pb.UserListResponse{}
	rsp.Total = int32(total)

	for _, user := range list {
		userInfoRsp := UserResponse(user)
		rsp.Data = append(rsp.Data, userInfoRsp)
	}

	return rsp, nil
}

func UserResponse(user *biz.User) *pb.UserInfoResponse {
	userInfoRsp := &pb.UserInfoResponse{
		Id:           user.ID,
		Email:        user.Email,
		Phone:        user.Phone,
		LastName:     user.Lastname,
		FirstName:    user.Firstname,
		FullName:     user.Firstname + " " + user.Lastname,
		CreatedAt:    timestamppb.New(user.CreatedAt),
		UpdatedAt:    timestamppb.New(user.UpdatedAt),
		Organization: user.Organization,
		Blocked:      user.Blocked,
		IsVerify:     user.IsVerify,
		Role:         user.Role,
		Messengers: &pb.Messenger{
			KakaoTalk: user.Messengers.KakaoTalk,
			Telegram:  user.Messengers.Telegram,
			WhatsApp:  user.Messengers.WhatsApp,
			Line:      user.Messengers.Line,
			Signal:    user.Messengers.Signal,
		},
	}
	return userInfoRsp
}

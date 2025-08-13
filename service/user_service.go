package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/TAhirr01/firstmodule/models"
	"github.com/TAhirr01/firstmodule/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type UserService struct {
	db *gorm.DB
	pb.UnimplementedUserServiceServer
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (service *UserService) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("UserService::RegisterUser called")

	// Eyni email ilə user var?
	var existing models.User
	if err := service.db.Where("email = ?", req.GetEmail()).First(&existing).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err == nil {
		return nil, ErrUserAlreadyExists
	}

	// Yeni user yarat
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := models.User{}
	newUser.SetName(req.GetName())
	newUser.SetEmail(req.GetEmail())
	newUser.SetPassword(string(hashedPassword))
	newUser.SetAge(req.GetAge())

	if err := service.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Name:  newUser.GetName(),
		Email: newUser.GetEmail(),
	}, nil
}

func (service *UserService) FindUserById(ctx context.Context, req *pb.UserId) (*pb.UserResponse, error) {
	fmt.Println("UserService::FindUserById called")
	var user *models.User
	err := service.db.First(&user, req.GetId()).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &pb.UserResponse{
		Id:    int32(user.ID),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil
}
func (service *UserService) FindAllUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetAllUsersResponse, error) {
	fmt.Println("UserService::FindAllUsers called")
	var users []models.User
	err := service.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	pbUsers := make([]*pb.UserResponse, len(users))
	for i, user := range users {
		pbUsers[i] = &pb.UserResponse{
			Id:    int32(user.ID),
			Name:  user.GetName(),
			Email: user.GetEmail(),
		}
	}
	return &pb.GetAllUsersResponse{
		Users: pbUsers,
	}, nil
}

func (service *UserService) FindUserByEmail(ctx context.Context, req *pb.Email) (*pb.UserResponse, error) {
	fmt.Println("UserService::FindUserByEmail called")
	var user *models.User
	err := service.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &pb.UserResponse{
		Id:           int32(user.ID),
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		PasswordHash: user.GetPassword(),
	}, nil

}

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/TAhirr01/firstmodule/models"
	"github.com/TAhirr01/firstmodule/pb"
	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type UserService struct {
	db *gorm.DB
	pb.UserServiceServer
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (service *UserService) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("UserService::RegisterUser called")
	var user *models.User
	err := service.db.Where("email = ?", req.GetEmail()).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	user = &models.User{
		Email:    req.GetEmail(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
		Age:      req.GetAge(),
	}
	err = service.db.Create(user).Error
	if err != nil {
		fmt.Println("cannot create user")
	}

	userDto := &pb.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return userDto, nil
}

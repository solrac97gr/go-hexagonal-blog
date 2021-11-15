package services

import (
	"errors"
	"goHexagonalBlog/internal/core/ports"
)

type UserService struct {
	userRepository ports.IUserRepository
}

//This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Register(email string, password string, confirmPass string) error {
	if password != confirmPass {
		return errors.New("the passwords are not equal")
	}
	err := s.userRepository.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}

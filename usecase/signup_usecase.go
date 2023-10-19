package usecase

import (
	"github.com/hitoshi-w/go-lang/model"
	"github.com/hitoshi-w/go-lang/repository"
	"github.com/hitoshi-w/go-lang/utility"
)

type SignupUsecase interface {
	Create(user *model.User) error
	GetByEmail(email string) (model.User, error)
	CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error)
}

type signupUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(ur repository.UserRepository) SignupUsecase {
	return &signupUsecase{
		userRepository: ur,
	}
}

func (su *signupUsecase) Create(user *model.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetByEmail(email string) (model.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return utility.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return utility.CreateRefreshToken(user, secret, expiry)
}

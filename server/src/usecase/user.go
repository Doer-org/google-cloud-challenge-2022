package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IUserUsecase interface {
	Create(ctx context.Context, age int, name string, authenticated bool, mail string, icon string) error
	GetByMail(ctx context.Context, mail string) (*entity.User, error)
	GetById(ctx context.Context, id string) (*entity.User, error)
}

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(r repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (u *UserUsecase) Create(ctx context.Context, age int, name string, authenticated bool, mail string, icon string) error {
	if name == "" {
		return fmt.Errorf("UserUsecase: name is empty")
	}
	if icon == "" {
		return fmt.Errorf("UserUsecase: icon is empty")
	}
	// TODO: mailが存在するかの確認

	user := &entity.User{
		Age:           age,
		Name:          name,
		Authenticated: authenticated,
		Mail:          mail,
		Icon:          icon,
	}
	return u.repo.Create(ctx, user)
}

func (u *UserUsecase) GetByMail(ctx context.Context, mail string) (*entity.User, error) {
	if mail == "" {
		return nil, fmt.Errorf("UserUsecase: mail is empty")
	}
	return u.repo.GetByMail(ctx, mail)
}

func (u *UserUsecase) GetById(ctx context.Context, id string) (*entity.User, error) {
	uId := entity.UserId(id)
	if uId == "" {
		return nil, fmt.Errorf("UserUsecase: id is empty")
	}
	return u.repo.GetById(ctx, uId)
}

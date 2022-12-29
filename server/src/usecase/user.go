package usecase

import (
	"context"
	"errors"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IUserUsecase interface {
	Create(ctx context.Context,age int,name string, authenticated bool,gmail string,icon_img string) error
}

type UserUsecase struct {
	Repo repository.IUserRepository
}

func NewUserUsecase(r repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		Repo: r,
	}
}

func (u *UserUsecase) Create(ctx context.Context,age int,name string, authenticated bool, gmail string,icon_img string) error {

	if age == 0 {
		return errors.New("age is not validated")
	}

	user := &entity.User{
		Age: age,
		Name: name,
		Authenticated: authenticated,
		Gmail: gmail,
		Icon_img: icon_img,
	}
	err := u.Repo.Create(ctx,user)
	return err
}

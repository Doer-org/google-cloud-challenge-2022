package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/service"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
)

type IUser interface {
	CreateNewUser(ctx context.Context, name string, mail string, icon string) (*ent.User, error)
	GetUserById(ctx context.Context, userIdString string) (*ent.User, error)
	DeleteUserById(ctx context.Context, userIdString string) error
	UpdateUserById(ctx context.Context, userIdString string, name string) (*ent.User, error)
	GetUserByMail(ctx context.Context, mail string) (*ent.User, error)
	GetUserEvents(ctx context.Context, userIdString string) ([]*ent.Event, error)
}

type User struct {
	repo repository.IUser
}

func NewUser(r repository.IUser) IUser {
	return &User{
		repo: r,
	}
}

func (uc *User) CreateNewUser(ctx context.Context, name string, mail string, icon string) (*ent.User, error) {
	found, err := uc.repo.GetUserByMail(ctx, mail)
	if err != nil {
		return nil, fmt.Errorf("getUserByMail: %w", err)
	}
	if found != nil {
		return nil, fmt.Errorf("this mail user is already exists")
	}
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	// iconがなかった場合,defaultのアイコンにする
	if icon == "" {
		icon = service.GetRandomDefaultIcon()
	}
	user := &ent.User{
		Name:          name,
		Authenticated: true,
		Mail:          mail,
		Icon:          icon,
	}
	return uc.repo.CreateNewUser(ctx, user)
}

func (uc *User) GetUserById(ctx context.Context, userIdString string) (*ent.User, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, fmt.Errorf("userId Parse: %w", err)
	}
	return uc.repo.GetUserById(ctx, userId)
}

func (uc *User) DeleteUserById(ctx context.Context, userIdString string) error {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return fmt.Errorf("userId Parse: %w", err)
	}
	return uc.repo.DeleteUserById(ctx, userId)
}

func (uc *User) UpdateUserById(ctx context.Context, userIdString string, name string) (*ent.User, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, fmt.Errorf("userId Parse: %w", err)
	}
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	// 変更できるのはnameのみ
	u := &ent.User{
		Name: name,
	}
	return uc.repo.UpdateUserById(ctx, userId, u)
}

func (uc *User) GetUserByMail(ctx context.Context, mail string) (*ent.User, error) {
	if mail == "" {
		return nil, fmt.Errorf("mail is empty")
	}
	return uc.repo.GetUserByMail(ctx, mail)
}

func (uc *User) GetUserEvents(ctx context.Context, userIdString string) ([]*ent.Event, error) {
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return nil, fmt.Errorf("userId Parse: %w", err)
	}
	return uc.repo.GetUserEvents(ctx, userId)
}

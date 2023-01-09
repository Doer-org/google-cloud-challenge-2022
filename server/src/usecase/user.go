package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/service"
)

type IUserUsecase interface {
	CreateNewUser(ctx context.Context, name string, authenticated bool, mail string, icon string) (*entity.User, error)
	GetUserById(ctx context.Context, userIdString string) (*entity.User, error)
	DeleteUserById(ctx context.Context, userIdString string) error
	UpdateUserById(ctx context.Context, userIdString string, name string, authenticated bool, mail string, icon string) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	GetEventAdminById(ctx context.Context,eventIdString string) (*entity.User,error)
}

type UserUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(r repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (uc *UserUsecase) CreateNewUser(ctx context.Context, name string, authenticated bool, mail string, icon string) (*entity.User, error) {
	if name == "" {
		return nil, fmt.Errorf("UserUsecase: name is empty")
	}
	// iconがなかった場合,defaultのアイコンにする
	if icon == "" {
		icon = service.GetRandomDefaultIcon()
	}
	// TODO: mailが存在するかの確認
	user := &entity.User{
		Name:          name,
		Authenticated: authenticated,
		Mail:          mail,
		Icon:          icon,
	}
	return uc.repo.CreateNewUser(ctx, user)
}

func (uc *UserUsecase) GetUserById(ctx context.Context, userIdString string) (*entity.User, error) {
	userId := entity.UserId(userIdString)
	if userId == "" {
		return nil, fmt.Errorf("UserUsecase: userId parse failed")
	}
	return uc.repo.GetUserById(ctx, userId)
}

func (uc *UserUsecase) DeleteUserById(ctx context.Context,userIdString string) error {
	userId := entity.UserId(userIdString)
	if userId == "" {
		return fmt.Errorf("UserUsecase: userId parse failed")
	}
	return uc.repo.DeleteUserById(ctx, userId)
}

func (uc *UserUsecase) UpdateUserById(ctx context.Context, userIdString string, name string, authenticated bool, mail string, icon string) (*entity.User, error) {
	userId := entity.UserId(userIdString)
	if userId == "" {
		return nil,fmt.Errorf("UserUsecase: userId parse failed")
	}
	if name == "" {
		return nil, fmt.Errorf("UserUsecase: name is empty")
	}
	// TODO: iconが空文字のときの処理を追加する
	// TODO: 更新できるのは本来認証済みユーザーのみ?
	u := &entity.User{
		Name: name,
		Authenticated: authenticated,
		Mail: mail,
		Icon: icon,
	}
	return uc.repo.UpdateUserById(ctx, userId,u)
}

func (uc *UserUsecase) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	if mail == "" {
		return nil, fmt.Errorf("UserUsecase: mail is empty")
	}
	return uc.repo.GetUserByMail(ctx, mail)
}

// TODO: これの場所ここで会ってる?
func (uc *UserUsecase) GetEventAdminById(ctx context.Context,eventIdString string) (*entity.User,error) {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("UserUsecase: eventId parse failed")
	}
	return uc.repo.GetEventAdminById(ctx,eventId)
}

package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(c *ent.Client) repository.IUserRepository {
	return &UserRepository{
		client: c,
	}
}

func (r *UserRepository) CreateNewUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	entUser, err := r.client.User.
		Create().
		SetName(u.Name).
		SetAuthenticated(u.Authenticated).
		SetMail(u.Mail).
		SetIcon(u.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: create user query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func (r *UserRepository) GetUserById(ctx context.Context, userId entity.UserId) (*entity.User, error) {
	userUuid, err := uuid.Parse(string(userId))
	if err != nil {
		return nil, fmt.Errorf("UserRepository: userUuid parse error: %w", err)
	}
	entUser, err := r.client.User.
		Query().
		Where(user.ID(userUuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: get user query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func (r *UserRepository) DeleteUserById(ctx context.Context, userId entity.UserId) error {
	userUuid, err := uuid.Parse(string(userId))
	if err != nil {
		return fmt.Errorf("UserRepository: userUuid parse error: %w", err)
	}
	err = r.client.User.
		DeleteOneID(userUuid).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("UserRepository: delete user query error: %w", err)
	}
	return nil
}

func (r *UserRepository) UpdateUserById(ctx context.Context, userId entity.UserId, u *entity.User) (*entity.User, error) {
	userUuid, err := uuid.Parse(string(userId))
	if err != nil {
		return nil,fmt.Errorf("UserRepository: userUuid parse error: %w", err)
	}
	entUser,err := r.client.User.
		UpdateOneID(userUuid).
		SetName(u.Name).
		SetAuthenticated(u.Authenticated).
		SetMail(u.Mail).
		SetIcon(u.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: update user query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func (r *UserRepository) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Mail(mail)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: get user query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func EntToEntityUser(e *ent.User) *entity.User {
	return &entity.User{
		Id:            entity.UserId(e.ID.String()),
		Name:          e.Name,
		Authenticated: e.Authenticated,
		Mail:          e.Mail,
		Icon:          e.Icon,
	}
}

func EntityToEntUser(u *entity.User) *ent.User {
	return &ent.User{
		Name:          u.Name,
		Authenticated: u.Authenticated,
		Mail:          u.Mail,
		Icon:          u.Icon,
	}
}

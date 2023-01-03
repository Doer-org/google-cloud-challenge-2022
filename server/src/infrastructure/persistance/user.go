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
	Client *ent.Client
}

func NewUserRepository(c *ent.Client) repository.IUserRepository {
	return &UserRepository{
		Client: c,
	}
}

func (r *UserRepository) Create(ctx context.Context, u *entity.User) error {
	_, err := r.Client.User.
		Create().
		SetAge(u.Age).
		SetName(u.Name).
		SetAuthenticated(u.Authenticated).
		SetMail(u.Mail).
		SetIcon(u.Icon).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("UserRepository: query error: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByMail(ctx context.Context, mail string) (*entity.User, error) {
	eu, err := r.Client.User.
		Query().
		Where(user.Mail(mail)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: query error: %w", err)
	}
	return entToEntity(eu), nil
}

func (r *UserRepository) GetById(ctx context.Context, id entity.UserId) (*entity.User, error) {
	uuid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, fmt.Errorf("UserRepository: uuid parse error: %w", err)
	}
	eu, err := r.Client.User.
		Query().
		Where(user.ID(uuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: query error: %w", err)
	}
	return entToEntity(eu), nil
}

func entToEntity(e *ent.User) *entity.User {
	return &entity.User{
		Id:            entity.UserId(e.ID.String()),
		Age:           e.Age,
		Name:          e.Name,
		Authenticated: e.Authenticated,
		Mail:          e.Mail,
		Icon:          e.Icon,
	}
}

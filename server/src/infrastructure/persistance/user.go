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
		return nil, fmt.Errorf("UserRepository: user create query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func (r *UserRepository) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	entUser, err := r.client.User.
		Query().
		Where(user.Mail(mail)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: user query error: %w", err)
	}
	return EntToEntityUser(entUser), nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id entity.UserId) (*entity.User, error) {
	userUuid, err := uuid.Parse(string(id))
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


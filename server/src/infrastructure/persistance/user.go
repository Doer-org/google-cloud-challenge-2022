package persistance

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type UserRepository struct {
	Client *ent.Client
}

func NewUserRepository(c *ent.Client) repository.IUserRepository {
	return &UserRepository{
		Client: c,
	}
}

func (r *UserRepository) Create(ctx context.Context,u *entity.User) error {
	_,err := r.Client.User.
				Create().
				SetAge(u.Age).
				SetName(u.Name).
				SetAuthenticated(u.Authenticated).
				SetGmail(u.Gmail).
				SetIconImg(u.Icon_img).
				Save(ctx)
	return err
}
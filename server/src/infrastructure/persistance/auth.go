package persistance

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type AuthRepository struct {
	Client *ent.Client
}

func NewAuthRepository(c *ent.Client) repository.IAuthRepository {
	return &AuthRepository{
		Client: c,
	}
}




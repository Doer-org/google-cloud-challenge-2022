package persistance

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type ImageRepository struct {
	Client *ent.Client
}

func NewImageRepository(c *ent.Client) repository.IImageRepository {
	return &ImageRepository{
		Client: c,
	}
}


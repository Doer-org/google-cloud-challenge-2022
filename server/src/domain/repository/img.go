package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IImageRepository interface {
	CreateImg(ctx context.Context, i *entity.Image) (*entity.Image, error)
}

package usecase

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	usecase_error "github.com/Doer-org/google-cloud-challenge-2022/error/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/hash"
)

type IImageUsecase interface {
	Create(ctx context.Context, img []byte) (*entity.Image, error)
}

type ImageUsecase struct {
	Repo repository.IImageRepository
}

func NewImageUsecase(r repository.IImageRepository) IImageUsecase {
	return &ImageUsecase{
		Repo: r,
	}
}

func (i *ImageUsecase) Create(ctx context.Context, img []byte) (*entity.Image, error) {
	if img == nil {
		return nil, usecase_error.ImgEmptyError
	}
	id := hash.GetUlid()

	image := &entity.Image{
		Id:  id,
		Img: img,
	}

	image, err := i.Repo.Create(ctx, image)
	return image, err
}

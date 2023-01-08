package service

import (
	"math/rand"
	"time"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

// TODO: Icon型を用意したほうが良い?
func GetRandomDefaultIcon() string {
	rand.Seed(time.Now().UnixNano())
	return entity.DEFAULT_ICON_LIST[rand.Intn(len(entity.DEFAULT_ICON_LIST))]
}

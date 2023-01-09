package service

import (
	"math/rand"
	"time"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
)

func GetRandomDefaultIcon() string {
	rand.Seed(time.Now().UnixNano())
	return constant.DEFAULT_ICON_LIST[rand.Intn(len(constant.DEFAULT_ICON_LIST))]
}

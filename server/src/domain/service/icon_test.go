package service

import (
	"testing"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
)

func Test_GetRandomDefaultIcon(t *testing.T) {
	str := GetRandomDefaultIcon()
	var correction bool
	correction = false
	for _, v := range constant.DEFAULT_ICON_LIST {
		if v == str {
			correction = true
		}
	}

	if correction == false {
		t.Errorf("get random icon faild")
	}
}

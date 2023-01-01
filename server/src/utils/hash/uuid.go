package hash

import (
	"github.com/google/uuid"
)

func GetUuid() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	uu := u.String()
	return uu, nil
}

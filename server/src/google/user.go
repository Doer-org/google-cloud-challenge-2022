package google

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
)

func (c *Client) GetMe(ctx context.Context) (*entity.User, error) {
	token, ok := utils.GetTokenFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	client := conf.Client(ctx, tok)
}

type GoogleUser struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

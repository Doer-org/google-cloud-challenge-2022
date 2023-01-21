package google

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"
)

func (c *Client) GetMe(ctx context.Context) (*ent.User, error) {
	token, ok := utils.GetTokenFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	// tokenを使用して、clientを返す
	client := c.auth.Config.Client(ctx, token)
	resp, err := client.Get(env.GetEnvOrDefault("GOOGLE_API_CLIENT",""))
	if err != nil {
		return nil, fmt.Errorf("googleapis Get: %w", err)
	}
	defer resp.Body.Close()
	var gUser googleUser
	if err:=json.NewDecoder(resp.Body).Decode(&gUser);err!=nil{
		return nil, fmt.Errorf("decode: %w", err)
	}
	user := &ent.User{
		Name:          gUser.Name,
		Mail:          gUser.Email,
		Icon:          gUser.Picture,
		Authenticated: true,
	}
	return user, nil
}

type googleUser struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

package google

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/config"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
)

func (c *Client) GetMe(ctx context.Context) (*ent.User, error) {
	token, ok := mycontext.GetToken(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	// tokenを使用して、clientを返す
	client := c.auth.Config.Client(ctx, token)
	resp, err := client.Get(config.GOOGLE_API_CLIENT)
	if err != nil {
		return nil, fmt.Errorf("googleapis Get: %w", err)
	}
	defer resp.Body.Close()
	var j googleUserJson
	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	user := &ent.User{
		Name:          j.Name,
		Mail:          j.Email,
		Icon:          j.Picture,
		Authenticated: true,
	}
	return user, nil
}

type googleUserJson struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

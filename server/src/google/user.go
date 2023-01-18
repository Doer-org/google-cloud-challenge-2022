package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/google/uuid"
)

func (c *Client) GetMe(ctx context.Context) (*ent.User, error) {
	token, ok := utils.GetTokenFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	client := c.auth.Config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("google api response get error: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("google api response read error: %w", err)
	}

	var ResUser GoogleUser
	if err := json.Unmarshal(body, &ResUser); err != nil {
		return nil, fmt.Errorf("google api json unmarshal error: %w", err)
	}

	userUuid, err := uuid.Parse(ResUser.Id)
	if err != nil {
		return nil, fmt.Errorf("uuid parse err : %w", err)
	}

	user := &ent.User{
		ID:            userUuid,
		Name:          ResUser.Name,
		Mail:          ResUser.Email,
		Icon:          ResUser.Picture,
		Authenticated: true,
	}
	return user, nil
}

type GoogleUser struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
)

func (c *Client) GetMe(ctx context.Context) (*ent.User, error) {
	token, ok := utils.GetTokenFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	// tokenを使用して、clientを返す
	client := c.auth.Config.Client(ctx, token)
	//TODO:環境変数にする
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("googleapis Get: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("googleapis response ReadAll: %w", err)
	}

	var ResUser GoogleUser
	if err := json.Unmarshal(body, &ResUser); err != nil {
		return nil, fmt.Errorf("googleapis user Unmarshal: %w", err)
	}
	//TODO: 消し忘れない
	log.Println(ResUser)
	user := &ent.User{
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

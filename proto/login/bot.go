package login

import (
	"context"

	"github.com/gotd/td/telegram"
)

func (LoginService) Bot(ctx context.Context, apiID int, apiHash, botToken string) (*telegram.Client, error) {
	client := telegram.NewClient(apiID, apiHash, telegram.Options{})
	err := client.Run(ctx, func(ctx context.Context) error {
		auth := client.Auth()
		_, err := auth.Bot(ctx, botToken)
		return err
	})
	return client, err
}

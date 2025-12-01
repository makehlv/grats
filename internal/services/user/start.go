package user

import (
	"context"
	"fmt"

	tgbot "github.com/makehlv/tgbot"
)

func (s *Service) Start(ctx context.Context, scope *tgbot.Scope) error {
	message := scope.Update().GetMessage()

	err := s.db.Tx(ctx, func(ctx context.Context) error {
		return s.userRegistration.RegisterOrUpdateUser(ctx, scope)
	})
	if err != nil {
		return err
	}

	username := message.From.Username
	if username == "" {
		username = message.From.FirstName
		if username == "" {
			username = s.cfg.Constants.GREETING_FRIEND
		}
	}

	hello := fmt.Sprintf(
		s.cfg.Constants.GREETING_TEMPLATE,
		username,
	)

	if _, err := scope.Reply(ctx, hello); err != nil {
		return err
	}

	return nil
}

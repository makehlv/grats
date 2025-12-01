package app

import (
	"context"
	"strings"

	tgbot "github.com/makehlv/tgbot"
)

func SupportReplyCondition(supportChatId string) tgbot.Condition {
	return func(ctx context.Context, scope *tgbot.Scope) (bool, error) {
		if scope.Update().GetMessage() != nil &&
			scope.Update().GetMessage().GetChatIdStr() == supportChatId &&
			scope.Update().GetMessage().IsReply() {
			return true, nil
		}

		return false, nil
	}
}

func ShowSharedListCondition() tgbot.Condition {
	return func(ctx context.Context, scope *tgbot.Scope) (bool, error) {
		command := scope.Update().GetMessage().GetCommand()

		if strings.HasPrefix(command, "/start") {
			idForCommand := strings.TrimSpace(strings.TrimPrefix(command, "/start"))

			if idForCommand != "" && strings.HasPrefix(idForCommand, "wl") {
				return true, nil
			}
		}

		return false, nil
	}
}

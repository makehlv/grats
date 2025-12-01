package wish

import (
	"context"
	"errors"

	"github.com/makehlv/grats/internal/repositories/models"
	"github.com/makehlv/grats/internal/repositories/wish_list"
	tgbot "github.com/makehlv/tgbot"
	inlinekeyboard "github.com/makehlv/tgbot/builders/inline_keyboard"
)

func (s *Service) PickFirstWishList(ctx context.Context, userId string) (*models.WishList, error) {
	filter := wish_list.ListFilter{UserId: userId}
	wishList, err := s.repositories.WishList.List(ctx, &filter)
	if err != nil {
		return nil, err
	}

	if len(wishList) == 0 {
		return nil, errors.New("wish list is empty")
	}

	return wishList[0], nil
}

func (s *Service) BuildEntityButtons(scope *tgbot.Scope, wishes []*models.Wish, offset int, callback func(id string, offset int) string) *inlinekeyboard.Builder {
	buttons := scope.Keyboard()
	for _, entity := range wishes {
		buttonText := entity.ButtonText()

		buttons.AppendAsLine(buttons.NewButton(buttonText, callback(entity.ID, offset)))
	}

	return buttons
}

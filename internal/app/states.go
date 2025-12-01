package app

import (
	"context"

	"github.com/makehlv/grats/internal/config"
	"github.com/makehlv/grats/internal/repositories"
	"github.com/makehlv/grats/internal/services"
	tgbot "github.com/makehlv/tgbot"
)

func RegisterHandlers(
	b *tgbot.Bot,
	s *services.Services,
	cfg *config.Config,
	repositories *repositories.Repositories,
) {
	c := cfg.Constants

	resetUserCache := func(ctx context.Context, scope *tgbot.Scope) error {
		return repositories.Cache.Reset(ctx, scope.Update().GetChatIdStr())
	}

	b.AddMiddleware(
		func(ctx context.Context, scope *tgbot.Scope) error {
			return scope.AnswerCallbackQuery(ctx)
		},
	)

	b.AddHandler(
		s.User.Start,
		tgbot.Command(c.CMD_START),
	)

	// ------------------------ support --------------------------------

	b.AddHandler(
		s.Support.Support,
		tgbot.Command(c.CMD_SUPPORT),
	)

	supportWrite := b.AddHandler(
		s.Support.SupportWrite,
		tgbot.CallbackDataContains(c.CMD_SUPPORT_WRITE),
		tgbot.BeforeAction(resetUserCache),
	)

	b.AddHandler(
		s.Support.SendSupportMessage,
		tgbot.MessageHasText(),
		tgbot.AcceptFrom(supportWrite),
	)

	b.AddHandler(
		s.Support.CancelSupportCall,
		tgbot.CallbackDataContains(c.CMD_SUPPORT_CANCEL),
	)

	b.AddHandler(
		s.Support.ProcessSupportReply,
		SupportReplyCondition(cfg.SupportChatId),
	)

	// ------------------------ wishlist --------------------------------

	addWish := b.AddHandler(
		s.Wish.AddWish,
		tgbot.CallbackDataContains(c.CMD_ADD_TO_WISH),
		tgbot.BeforeAction(resetUserCache),
	)

	b.AddHandler(
		s.Wish.SaveWish,
		tgbot.MessageHasText(),
		tgbot.AcceptFrom(addWish),
	)

	b.AddHandler(
		s.Wish.List,
		tgbot.Command(c.CMD_WISHLIST),
	)

	b.AddHandler(
		s.Wish.List,
		tgbot.CallbackDataContains(c.CMD_LIST),
	)

	b.AddHandler(
		s.Wish.WishInfo,
		tgbot.CallbackDataContains(c.CMD_WISH_INFO),
	)

	b.AddHandler(
		s.Wish.DeleteWish,
		tgbot.CallbackDataContains(c.CMD_DELETE_WISH),
	)

	b.AddHandler(
		s.Wish.ConfirmDeleteWish,
		tgbot.CallbackDataContains(c.CMD_CONFIRM_DELETE_WISH),
	)

	editWishName := b.AddHandler(
		s.Wish.EditWishName,
		tgbot.CallbackDataContains(c.CMD_EDIT_WISH_NAME),
		tgbot.BeforeAction(resetUserCache),
	)

	b.AddHandler(
		s.Wish.SaveEditWishName,
		tgbot.MessageHasText(),
		tgbot.AcceptFrom(editWishName),
	)

	editWishLink := b.AddHandler(
		s.Wish.EditLink,
		tgbot.CallbackDataContains(c.CMD_EDIT_LINK),
		tgbot.BeforeAction(resetUserCache),
	)

	b.AddHandler(
		s.Wish.SaveEditLink,
		tgbot.MessageHasText(),
		tgbot.AcceptFrom(editWishLink),
	)

	editWishPrice := b.AddHandler(
		s.Wish.EditPrice,
		tgbot.CallbackDataContains(c.CMD_EDIT_PRICE),
		tgbot.BeforeAction(resetUserCache),
	)

	b.AddHandler(
		s.Wish.SaveEditPrice,
		tgbot.MessageHasText(),
		tgbot.AcceptFrom(editWishPrice),
	)

	b.AddHandler(
		s.Wish.DeleteLink,
		tgbot.CallbackDataContains(c.CMD_DELETE_LINK),
		tgbot.AcceptFrom(editWishLink),
	)

	b.AddHandler(
		s.Wish.ShareWishList,
		tgbot.CallbackDataContains(c.CMD_SHARE_WISH_LIST),
	)

	b.AddHandler(
		s.Wish.ToggleWishLock,
		tgbot.CallbackDataContains(c.CMD_TOGGLE_WISH_LOCK),
	)

	b.AddHandler(
		s.Wish.ShowSharedWishlist,
		ShowSharedListCondition(),
	)

	b.AddHandler(
		s.Wish.WishInfo,
		tgbot.CallbackDataContains(c.CMD_SHOW_SWI),
	)

	b.AddHandler(
		s.Wish.ShowSharedWishlist,
		tgbot.CallbackDataContains(c.CMD_SHOW_SWL),
	)

	b.Reset(
		tgbot.Command(c.CMD_CANCEL),
		tgbot.BeforeAction(resetUserCache),
		tgbot.AcceptFrom(supportWrite),
		tgbot.AcceptFrom(editWishName),
		tgbot.AcceptFrom(editWishLink),
		tgbot.AcceptFrom(editWishPrice),
		tgbot.AcceptFrom(addWish),
	)
}

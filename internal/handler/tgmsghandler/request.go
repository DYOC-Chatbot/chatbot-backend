package tgmsghandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	autherror "backend/pkg/error/externalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	RequestCmdWord = "request"
	RequestCmdDesc = "Make a logistical request"
)

const (
	AuthRequiredErrorResponse = "Authentication required before request can be processed, please provide your booking ID"
)

func HandleRequestCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		return NoChatFoundResponse, err
	}

	bk, _ := booking.ReadByChatID(db, chat.ID)

	if err := createRequestQueryTransaction(db, msg, chat, bk, model.TypeRequest); err != nil {
		if autherror.IsAuthRequiredError(err) {
			return AuthRequiredErrorResponse, err
		}
		return "An error occurred while creating a new request", err
	}

	return "New query created", nil
}

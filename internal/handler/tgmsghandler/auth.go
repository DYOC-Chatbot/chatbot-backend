package tgmsghandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/viewmodel"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"
	"encoding/json"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AuthCmdWord = "auth"
	AuthCmdDesc = "Authenticate yourself to make a request. Please provide your credentials"
)

const (
	AuthRequestMadeResponse = "Authentication request made. Pending response from staff."
)

var (
	CredentialsNotFound = errors.New("Credentials not found")
)

func HandleAuthCommand(msg *tgbotapi.Message, hub *ws.Hub) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return NoChatFoundResponse, nil
		}
		return "", err
	}

	cred, err := extractCredentials(msg.Text)
	if err != nil {
		return err.Error(), err
	}

	tgAuthView := viewmodel.TgAuthView{
		ChatID:      chat.ID,
		Credentials: cred,
	}

	msgStruct := api.WebSocketMessage{
		Type: api.AuthType,
		Data: tgAuthView,
	}

	msgBytes, err := json.Marshal(msgStruct)
	if err != nil {
		return "", err
	}

	hub.Broadcast <- msgBytes
	return AuthRequestMadeResponse, nil
}

// Commands are prefixed with a slash (/cmd args)
func extractCredentials(text string) (string, error) {
	for i, c := range text {
		if c == ' ' {
			if len(text) <= i+1 {
				return "", CredentialsNotFound
			}
			return text[i+1:], nil
		}
	}

	return "", CredentialsNotFound
}

package bothandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/handler/tgmessagehandler"
	"backend/internal/model"
	"backend/internal/params/tgmessageparams"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SendMessage(c echo.Context) error {
	chatIDStr := c.Param("chat_id")
	chatID, err := strconv.Atoi(chatIDStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid chat_id "+err.Error())
	}
	r := new(tgmessageparams.MessageParams)
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := database.GetDb()

	chat, err := chat.Read(db, uint(chatID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	msg, err := SendTelegramMessage(int64(chat.TelegramChatId), nil, r.Message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Assumption: This endpoint will only be used by staff when sending messages directly to the user
	msgModel, err := tgmessagehandler.SaveTgMessageToDB(db, msg, model.ByStaff)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Message sent: %s", msgModel.MessageBody))
}

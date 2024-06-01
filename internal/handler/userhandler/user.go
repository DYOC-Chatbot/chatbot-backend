package userhandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/user"
	"backend/internal/database"
	"backend/internal/viewmodel"
	"github.com/labstack/echo/v4"
	"net/http"
)

func List(c echo.Context) error {
	db := database.GetDb()
	users := user.ReadAll(db)
	userViews := make([]*viewmodel.UserView, len(users))
	for i, u := range users {
		userViews[i] = u.ToView()
	}
	c.Logger().Debugf("users: %v\n", userViews)
	return c.JSON(http.StatusOK, api.Response{Data: userViews})
}

func Create(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

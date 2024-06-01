package router

import (
	"backend/internal/handler/userhandler"
	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group) {
	g.GET("", userhandler.List)
}

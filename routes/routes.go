package routes

import (
	"app/handlers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// View routes
	e.GET("/", handlers.HomeHandler)

	e.GET("/services", handlers.ServicesHandler)

	e.GET("/contact", handlers.ContactHandler)
	e.POST("/contact", handlers.ContactSubmitHandler)

	e.GET("/about", handlers.AboutHandler)

	e.GET("/shipping", handlers.ShippingHandler)

	// Utility routes
	e.GET("/404", handlers.ErrorPageHandler)
	e.DELETE("/toast/remove", handlers.ToastRemover)
}

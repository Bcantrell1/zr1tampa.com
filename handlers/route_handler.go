package handlers

import (
	"app/templates/components"
	"app/views"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	components.Layout(views.Index()).Render(c.Request().Context(), c.Response())
	return nil
}

func ServicesHandler(c echo.Context) error {
	components.Layout(views.Services()).Render(c.Request().Context(), c.Response())
	return nil
}

func AboutHandler(c echo.Context) error {
	components.Layout(views.About()).Render(c.Request().Context(), c.Response())
	return nil
}

func ShippingHandler(c echo.Context) error {
	components.Layout(views.Shipping()).Render(c.Request().Context(), c.Response())
	return nil
}

func ContactHandler(c echo.Context) error {
	components.Layout(views.Contact()).Render(c.Request().Context(), c.Response())
	return nil
}

func ErrorPageHandler(c echo.Context) error {
	components.Layout(views.Error()).Render(c.Request().Context(), c.Response())
	return nil
}

func ContactSubmitHandler(c echo.Context) error {
	// Get form data
	contact := make(map[string]string)
	if err := c.Bind(&contact); err != nil {
		return err
	}

	// Get Fields
	first := contact["first_name"]
	last := contact["last_name"]
	email := contact["email"]
	phone := contact["phone"]
	bike_type := contact["bike_make"]
	bike_model := contact["model"]
	weight := contact["weight"]
	height := contact["height"]
	rider_skill := contact["rider_skill"]
	rider_style := contact["rider_style"]
	message := contact["message"]

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	to := "to@example.com"
	body := []byte("To: " + to + "\r\n" +

		"Subject: ZR1 Tampa Form Submission From " + first + " " + last + "\r\n" +

		"\r\n" +

		"Name: " + first + " " + last + "\r\n" +
		"Email: " + email + "\r\n" +
		"Phone: " + phone + "\r\n" + "\r\n" +
		"Bike: " + bike_type + " " + bike_model + "\r\n" + "\r\n" +
		"Rider Weight: " + weight + "\r\n" +
		"Rider Height: " + height + "\r\n" +
		"Rider Skill: " + rider_skill + "\r\n" +
		"Riding Style: " + rider_style + "\r\n" + "\r\n" +
		"Message: " + message + "\r\n")

	auth := smtp.PlainAuth("", username, password, "sandbox.smtp.mailtrap.io")

	err := smtp.SendMail("sandbox.smtp.mailtrap.io:587", auth, email, []string{to}, body)
	if err != nil {
		return components.Toast(false).Render(c.Request().Context(), c.Response())
	}

	// Send response
	return components.Toast(true).Render(c.Request().Context(), c.Response())
}

func ToastRemover(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

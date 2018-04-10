package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Giftcard struct {
	ID   int       `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Sum  int       `json:"sum"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	api := e.Group("/v1")
	giftcards := api.Group("/giftcards")
	giftcards.POST("/", createGiftCard)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createGiftCard(c echo.Context) error {
	giftcard := &Giftcard{
		UUID: uuid.New(),
	}
	if err := c.Bind(giftcard); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, giftcard)
}

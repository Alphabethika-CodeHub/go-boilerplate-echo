package main

import (
	"github.com/Alphabethika-CodeHub/go-boilerplate-echo/src/internal/utilities"
	"github.com/labstack/echo/v4"
)

func main() {
	// Connections
	utilities.LoadViper()

	// Initiate Server
	e := echo.New()

	// Middlewares

	// Routes

	// Start Server
	e.Logger.Fatal(e.Start(":3400"))
}

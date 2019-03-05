package main

import (
	"github.com/kardianos/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

type Prog struct {
}

func (Prog) Start(s service.Service) error {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handleReq)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func handleReq(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (Prog) Stop(s service.Service) error {
	println("Stopping service")
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}
	prg := &Prog{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Install()
	if err != nil {
		panic(err)
	}

}

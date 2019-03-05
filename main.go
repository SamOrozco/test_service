package main

import (
	"github.com/kardianos/service"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Prog struct {
}

func (prog Prog) Start(s service.Service) error {
	// Echo instance
	println("Starting this service")
	e := echo.New()
	// Routes
	e.GET("/", handleReq)
	go func() {
		e.Logger.Fatal(e.Start(":9008"))
	}()
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
		Name:        "TestService66",
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

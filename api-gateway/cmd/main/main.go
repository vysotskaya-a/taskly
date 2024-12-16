package main

import (
	"chat/internal/service"
	"chat/internal/ws"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	e := echo.New()
	hub := ws.NewHub(service.NewChat(rdb))
	handler := ws.NewHandler(hub)
	handler.Route(e)
	go hub.Run()
	port, _ := os.LookupEnv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

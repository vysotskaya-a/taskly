package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //TODO нужно будет убрать
}

type Handler struct {
	Hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

func (h *Handler) JoinRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		client := &Client{
			ID:      c.QueryParam("id"), //TODO будем брать из токена
			Conn:    conn,
			Updates: make(chan interface{}),
			RoomIDs: []string{"project-1", "project-2"}, //TODO будем брать чаты, в которых есть пользователь из Chat service
			Email:   c.QueryParam("email"),
		}
		h.Hub.Register <- client
		go client.ReadMessage(h.Hub)
		go client.WriteUpdates()
		return nil
	}
}

func (h *Handler) Route(e *echo.Echo) {
	e.GET("/ws", h.JoinRooms())
}

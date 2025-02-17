package main

import (
	"github.com/lxzan/gws"
	"log"
)

func main() {
	var app = gws.NewServer(new(Handler), &gws.ServerOption{
		CompressEnabled:  true,
		CheckUtf8Enabled: true,
	})
	log.Fatalf("%v", app.Run(":8000"))
}

type Handler struct {
	gws.BuiltinEventHandler
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	_ = socket.WriteMessage(message.Opcode, message.Bytes())
}

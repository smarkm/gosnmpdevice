package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/smarkm/gosnmpdevice/comps"
	g "github.com/soniah/gosnmp"
)

func main() {
	e := echo.New()
	e.Static("/ui", "ui")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", hello)
	StartTrapServer()
	e.Logger.Fatal(e.Start(":7780"))

}

func StartTrapServer() {
	server := comps.NewTrapServer()
	server.TrapHandler = myTrapHandler
	server.Start()

}

func myTrapHandler(packet *g.SnmpPacket, addr *net.UDPAddr) {
	log.Printf("got trapdata from %s\n", addr.IP)
	traps <- packet
}

var (
	upgrader = websocket.Upgrader{}
)
var traps = make(chan *g.SnmpPacket, 10)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	tag := make(chan int)
	ws.SetCloseHandler(func(code int, text string) error {
		tag <- 1
		return nil
	})
	for {
		select {
		case pkg := <-traps:
			err = ws.WriteJSON(pkg)
			if err != nil {
				println("err:", err)
			}
			println("get packet:", pkg.AgentAddress)
		case <-tag:
			break
		}
	}
	// Write

	// Read
	// _, msg, err := ws.ReadMessage()
	// if err != nil {
	// 	c.Logger().Error(err)
	// }
	// fmt.Printf("%s\n", msg)
	return err
}

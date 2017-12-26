package comps

import (
	"log"
	"net"
	"os"
	"strconv"

	g "github.com/soniah/gosnmp"
)

var Tport = 7162

type TrapServer struct {
	Tport       int
	TrapHandler func(s *g.SnmpPacket, u *net.UDPAddr)
}

func NewTrapServer() *TrapServer {
	return &TrapServer{Tport: Tport}
}
func (s *TrapServer) Start() {
	go s.run()
}

func (s *TrapServer) run() {
	tl := g.NewTrapListener()
	tl.OnNewTrap = s.TrapHandler
	tl.Params = g.Default
	tl.Params.Logger = log.New(os.Stdout, "", 0)
	log.Printf("TrapServer: listen on port %d", s.Tport)
	err := tl.Listen("0.0.0.0:" + strconv.Itoa(s.Tport))
	if err != nil {
		log.Panicf("error in listen: %s", err)
	}
}

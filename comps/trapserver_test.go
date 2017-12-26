package comps_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/smarkm/gosnmpdevice/comps"
)

func TestTrapServer(t *testing.T) {
	server := comps.NewTrapServer()
	server.Start()
	udpAddr, err := net.ResolveUDPAddr("", "0.0.0.0:"+strconv.Itoa(server.Tport))
	if err != nil {
		t.Fatal("Error: \n", err, udpAddr)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		t.Fatal("Error:", err, conn)
	}
}

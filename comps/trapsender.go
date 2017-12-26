package comps

import (
	"log"
	"os"

	g "github.com/soniah/gosnmp"
)

//DefaultTrapSender default sender
var DefaultTrapSender = &TrapSender{}

//TrapSender sender object
type TrapSender struct {
}

// SNMPConfig snmp params
type SNMPConfig g.GoSNMP

// Send send trap
func (s *TrapSender) Send(cfg g.GoSNMP, trap g.SnmpTrap) {
	g.Default.Target = cfg.Target
	g.Default.Port = cfg.Port
	g.Default.Version = cfg.Version
	g.Default.Community = cfg.Community
	g.Default.Logger = log.New(os.Stdout, "", 0)

	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	_, err = g.Default.SendTrap(trap)
	if err != nil {
		log.Fatalf("SendTrap() err: %v", err)
	}
}

func main() {

	// Default is a pointer to a GoSNMP struct that contains sensible defaults

}

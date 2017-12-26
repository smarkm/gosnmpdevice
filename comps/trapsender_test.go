package comps_test

import (
	"testing"

	"github.com/smarkm/gosnmpdevice/comps"

	g "github.com/soniah/gosnmp"
)

func TestSendTrap(t *testing.T) {
	sConfig := g.GoSNMP{Version: g.Version2c, Port: 7162}
	pdu := g.SnmpPDU{
		Name:  ".1.3.6.1.6.3.1.1.4.1.0",
		Type:  g.ObjectIdentifier,
		Value: ".1.3.6.1.6.3.1.1.5.1",
	}

	trap := g.SnmpTrap{
		Variables: []g.SnmpPDU{pdu},
	}
	comps.DefaultTrapSender.Send(sConfig, trap)
}

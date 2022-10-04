package NetSec

import (
	"testing"

	"github.com/reiver/go-telnet"
)

func TestEcho(t *testing.T) {
	Echo()
	// test not running the telnet client as expexted
	var caller telnet.Caller = telnet.StandardCaller
	telnet.DialToAndCall("0.0.0.0:3333", caller)
}

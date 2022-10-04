package NetSec

import (
	"fmt"
	"testing"
)

func TestCheckPorts(t *testing.T) {
	device := Device{}

	device.initDevice("scanme.nmap.org", "tcp")
	device.CheckPorts()
	for _, port := range device.portsOpen {
		fmt.Printf("%d open\n", port)
	}
}

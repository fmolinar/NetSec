// Testing for Port Availability

package NetSec

import (
	"fmt"
	"net"
	"sort"
)

type Device struct {
	portsChan   chan int
	host        string
	portsOpen   []int
	proto       string
	portsResult chan int
}

func (dv *Device) initDevice(host string, protocol string) *Device {
	// init
	dv.portsChan = make(chan int, 1000)
	dv.portsResult = make(chan int)
	dv.host = host
	dv.proto = protocol

	return dv
}

func (dv *Device) CheckPorts() {

	// start scanning the port
	for i := 0; i < cap(dv.portsChan); i++ {
		go dv.worker()
	}

	// traverse over the first 1024 ports
	go func() {
		for i := 1; i <= 1024; i++ {
			dv.portsChan <- i
		}
	}()

	// gather the results
	for i := 0; i < 1024; i++ {
		port := <-dv.portsResult
		if port != 0 {
			dv.portsOpen = append(dv.portsOpen, port)
		}
	}
	close(dv.portsResult)
	close(dv.portsChan)
	// sort the result
	sort.Ints(dv.portsOpen)

}

func (dv *Device) worker() {
	// ports := make(chan int, 100)
	for p := range dv.portsChan {
		address := fmt.Sprintf(dv.host+":%d", p)
		conn, err := net.Dial(dv.proto, address)
		if err != nil {
			dv.portsResult <- 0
			continue
		}
		conn.Close()
		dv.portsResult <- p
	}
}

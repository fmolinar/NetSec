package NetSec

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func ReadWrite(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(req), req)

	log.Println("Writing data back")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString("echo: " + req); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func Echo() {
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal("Unable to start tcp port 3333")
	}
	fmt.Println("Listenning on 0.0.0.0:3333")

	for {
		conn, err := listener.Accept()
		fmt.Println("Received connection on port 3333")
		if err != nil {
			log.Fatal("Unable to accept connection")
		}
		go ReadWrite(conn)
	}
}

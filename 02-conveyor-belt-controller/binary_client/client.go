package main

import (
	"fmt"
	"net"
	"os"
)

const (
	port = ":9001"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2)

	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		command := buffer[0]
		switch command {
		case 0x01:
			fmt.Println("Received Start Conveyor command")
			conn.Write([]byte{0x00})
		case 0x02:
			fmt.Println("Received Stop Conveyor command")
			conn.Write([]byte{0x00})
		case 0x03:
			speed := buffer[1]
			fmt.Printf("Received Set Speed command: %d\n", speed)
			conn.Write([]byte{0x00})
		default:
			fmt.Println("Unknown command")
			conn.Write([]byte{0xFF})
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Binary Protocol Conveyor Belt Client listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

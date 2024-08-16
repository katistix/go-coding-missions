package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	port = ":9002"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		command := scanner.Text()
		switch {
		case command == "START":
			fmt.Println("Received Start Conveyor command")
			conn.Write([]byte("OK\n"))
		case command == "STOP":
			fmt.Println("Received Stop Conveyor command")
			conn.Write([]byte("OK\n"))
		case strings.HasPrefix(command, "SPEED"):
			speed := strings.TrimPrefix(command, "SPEED ")
			fmt.Printf("Received Set Speed command: %s\n", speed)
			conn.Write([]byte("OK\n"))
		default:
			fmt.Println("Unknown command")
			conn.Write([]byte("ERROR Unknown command\n"))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Text-Based Protocol Conveyor Belt Client listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

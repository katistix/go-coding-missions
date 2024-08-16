package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	port = ":9003"
)

type Command struct {
	Command string `json:"command"`
	Value   int    `json:"value,omitempty"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	for {
		var cmd Command
		if err := decoder.Decode(&cmd); err != nil {
			fmt.Println("Error decoding JSON:", err)
			break
		}

		var resp Response

		switch cmd.Command {
		case "START":
			fmt.Println("Received Start Conveyor command")
			resp = Response{Status: "OK"}
		case "STOP":
			fmt.Println("Received Stop Conveyor command")
			resp = Response{Status: "OK"}
		case "SPEED":
			fmt.Printf("Received Set Speed command: %d\n", cmd.Value)
			resp = Response{Status: "OK"}
		default:
			fmt.Println("Unknown command")
			resp = Response{Status: "ERROR", Message: "Unknown command"}
		}

		if err := encoder.Encode(resp); err != nil {
			fmt.Println("Error encoding JSON:", err)
			break
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

	fmt.Println("JSON-Based Protocol Conveyor Belt Client listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

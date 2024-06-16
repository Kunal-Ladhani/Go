package main

import (
	"fmt"
	"net"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Read request
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("\nError reading request:", err)
		return
	}

	// Simulate processing the request (replace with actual logic)
	fmt.Printf("\nReceived request:\n%s\n", string(buf[:n]))

	// Send response
	response := "HTTP/1.1 200 OK\nContent-Type: text/plain\n\nHello, world!\n"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("\nError sending response:", err)
	}
}

func main() {
	const PORT = "8080"

	// Listen on port 8080
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Println("\nError listening:", err)
		return
	}
	defer ln.Close()

	fmt.Printf("\nServer listening on port %s", PORT)

	// Accept incoming connections in a loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("\nError accepting connection:", err)
			continue
		}
		go handleRequest(conn) // Handle connection concurrently using a goroutine
	}
}

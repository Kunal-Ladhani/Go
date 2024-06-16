package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP/1.1 200 OK\nContent-Type: text/plain\n\nHello, world!\n")
}

func main() {
	http.HandleFunc("/", handler)
	const PORT = "8080"

	fmt.Printf("Server listening on port: %s\n", PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		panic(err)
	}
}

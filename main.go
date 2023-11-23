package main

import (
	"fmt"
	"net/http"
)

const port string = ":8080"

func main() {
	fmt.Println("Hello from this side")

	http.HandleFunc("/", handler)

	fmt.Printf("Listening at port %s ", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from the server")
}

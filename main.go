package main

import (
	"net/http"

	"github.com/Spirolina/go-auth/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Welcome)
	http.ListenAndServe(":5500", nil)
}

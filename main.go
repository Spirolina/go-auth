package main

import (
	"log"
	"net/http"

	"github.com/Spirolina/go-auth/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

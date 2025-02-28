package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jeetnation/todoApp/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on Port 9000")

	log.Fatal(http.ListenAndServer(":9000", r))
}

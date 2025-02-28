package main

impor{
	"fmt"
	"log"
	"net/http"
	"GOLANG-React-TODO/router"
}


func main(){
	r := router.Router()
	fmt.Println("starting the server on Port 9000")

	log.Fatal(http.ListenAndServer(":9000", r))
}

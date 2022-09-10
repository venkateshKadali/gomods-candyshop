package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/venkateshKadali/gomods-candyshop.git/pkg/http/rest"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}

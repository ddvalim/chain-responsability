package main

import (
	"chain-responsability/config"
	"chain-responsability/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvVariables()

	r := router.CreateRouter()

	fmt.Printf("Server running on port %d\n", config.APIPort)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPort), r))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp-devbook/src/router"
)

func main() {

	fmt.Println("Rodando Web App")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}

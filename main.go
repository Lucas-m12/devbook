package main

import (
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	route := router.Generate()
	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", route))
}

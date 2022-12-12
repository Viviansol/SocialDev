package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/src/router"
)

func main() {

	fmt.Println("rodando")
	r := router.Generate()
	log.Fatalln(http.ListenAndServe(":3000", r))

}

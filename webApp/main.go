package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/src/router"
	"webApp/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()
	fmt.Println("rodando")

	log.Fatalln(http.ListenAndServe(":3080", r))

}

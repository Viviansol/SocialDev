package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/src/config"
	"webApp/src/router"
	"webApp/src/utils"
)

func main() {
	config.LoadConfig()
	utils.LoadTemplates()
	r := router.Generate()
	fmt.Println("rodando")

	log.Fatalln(http.ListenAndServe(":3080", r))

}

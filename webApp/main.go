package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/src/config"
	"webApp/src/cookies"
	"webApp/src/router"
	"webApp/src/utils"
)

func main() {
	config.LoadConfig()
	cookies.Configure()
	utils.LoadTemplates()
	r := router.Generate()
	fmt.Println("rodando")

	log.Fatalln(http.ListenAndServe(":3080", r))

}

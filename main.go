package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	fmt.Printf("escutando na porta %d", config.Porta)
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":5300", r))
}

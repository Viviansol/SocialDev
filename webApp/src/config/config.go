package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	APIURL   = " "
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

func LoadConfig() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	Porta, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

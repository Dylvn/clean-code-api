package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/Dylvn/dragonfire-api/src/.env"))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Dotenv loaded.")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

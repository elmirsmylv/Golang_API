package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoUri() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoUri := os.Getenv("MONGO_URI")

	return mongoUri

}

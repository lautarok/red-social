package infra

import "github.com/joho/godotenv"

func InitConfig() {
	godotenv.Load(".env")
}

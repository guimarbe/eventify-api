package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load("application.env")
	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file")
	}
}
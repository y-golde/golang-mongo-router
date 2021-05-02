package main

import (
	"helloworld/src/API"
	"helloworld/src/DB"

	"github.com/joho/godotenv"
)

func Hello() string {
	return "hello world"
}

func main() {
	godotenv.Load(".env")

	DB.InitClient()
	API.InitRouter()
}

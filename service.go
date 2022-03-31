package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	RABBITMQ_HOST  = ""
	RABBITMQ_USER  = ""
	RABBITMQ_PASS  = ""
	RABBITMQ_VHOST = ""
)

func init() {
	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
	RABBITMQ_VHOST = os.Getenv("VIRTUAL_HOST")
}

func main() {

}

func CreateNewUser(username string) {

}

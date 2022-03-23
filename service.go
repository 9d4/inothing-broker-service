package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var RABBITMQ_HOST = ""
var RABBITMQ_USER = ""
var RABBITMQ_PASS = ""

func init() {
	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
}

func main() {

}

package main

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
)

type part string

const (
	MQTT part = "MQTT    >> "
)

var (
	RABBITMQ_HOST  = ""
	RABBITMQ_USER  = ""
	RABBITMQ_PASS  = ""
	RABBITMQ_VHOST = ""
	MQTT_CLIENT_ID = "the-inothing-service"
	MQTT_SERVER    = ""
)

var mqttClient mqtt.Client

func init() {
	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
	RABBITMQ_VHOST = os.Getenv("VIRTUAL_HOST")
	MQTT_SERVER = os.Getenv("MQTT_SERVER")
}

func main() {
	mqttClient = runMqttClient()

	select {}
}

func runMqttClient() (client mqtt.Client) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(MQTT_SERVER)
	opts.SetClientID(MQTT_CLIENT_ID)
	opts.SetUsername(RABBITMQ_USER)
	opts.SetPassword(RABBITMQ_PASS)
	opts.SetCleanSession(true)
	opts.SetResumeSubs(true)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println(MQTT, "Connected!")
	return
}

func CreateNewUser(username string) {

}

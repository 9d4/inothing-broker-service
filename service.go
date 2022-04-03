package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	c := runMqttClient()
	tick := time.Tick(time.Second)

	for {
		<-tick
		fmt.Println(c.IsConnected())
	}

	select {}
}

var (
	MQTT_CLIENT_ID = "the-inothing-service"
	MQTT_SERVER    = os.Getenv("MQTT_SERVER")
)

func mqttMessageHandler(c mqtt.Client, m mqtt.Message) {
	fmt.Println("HIT")
}

func runMqttClient() (client mqtt.Client) {
	mqtt.DEBUG = log.New(os.Stdout, "MQTT:", log.Flags())

	opts := mqtt.NewClientOptions()
	opts.AddBroker(MQTT_SERVER)
	opts.SetClientID(MQTT_CLIENT_ID)
	opts.SetUsername(RABBITMQ_USER)
	opts.SetPassword(RABBITMQ_PASS)
	opts.SetDefaultPublishHandler(mqttMessageHandler)

	client = mqtt.NewClient(opts)
	if token:= client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	if true {
		if t := client.Subscribe("#", 0, nil); t.Wait() && t.Error() != nil {
			log.Println(t.Error())
			panic("Can't subscribe")
		}
	}

	return
}

func CreateNewUser(username string) {

}

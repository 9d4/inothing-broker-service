package main

import (
	"log"
	"net"
	"os"

	"github.com/9d4/inothing-broker-service/pb/userpb"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type part string

const (
	MQTT part = "MQTT    >> "
	GRPC part = "GRPC    >> "
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
var grpcListener net.Listener
var grpcServer *grpc.Server

func init() {
	RABBITMQ_HOST = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_USER = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASS = os.Getenv("RABBITMQ_PASS")
	RABBITMQ_VHOST = os.Getenv("VIRTUAL_HOST")
	MQTT_SERVER = os.Getenv("MQTT_SERVER")
}

func main() {
	grpcListener, grpcServer := startGrpcServer()
	mqttClient = startMqttClient()

	_, _ = grpcListener, grpcServer
	select {}
}

func startGrpcServer() (net.Listener, *grpc.Server) {
	listener, err := net.Listen("tcp", ":9200")
	if err != nil {
		log.Fatal("Can't start listener for grpc server. reason:", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)

	userpb.RegisterUserServiceServer(server, &GrpcServer{})

	go func() {
	restart:
		if err = server.Serve(listener); err != nil {
			log.Fatal("Can't serve grpc server. reason:", err)
			goto restart
		}
	}()
	log.Println(GRPC, "Up!")
	return listener, server
}

func startMqttClient() (client mqtt.Client) {
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

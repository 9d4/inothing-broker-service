package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCreateNewUser(t *testing.T) {
	if _, err := ReqCreateNewUser("user0", "user0"); err != nil {
		t.Fatal("error: ", err)
	}
}

func TestReqModVhost(t *testing.T) {
	path := fmt.Sprintf(ApiPermissionsVhostUser, normalizeNames(RABBITMQ_VHOST), "user0")
	res, err := Req("GET", path, nil)
	if err != nil {
		t.Error(err)
	}

	resStr := ResponseReadString(res)
	mod := VhostMod{}

	json.Unmarshal([]byte(resStr), &mod)
	fmt.Printf("%+v\n", mod)
	fmt.Println(res.StatusCode)
	fmt.Println(resStr)
}

func TestReqModTopic(t *testing.T) {
	path := fmt.Sprintf(ApiTopic_permissionsVhostUser, normalizeNames(RABBITMQ_VHOST), "user0")
	res, err := Req("GET", path, nil)
	if err != nil {
		t.Error(err)
	}

	resStr := ResponseReadString(res)
	mod := []TopicMod{}

	json.Unmarshal([]byte(resStr), &mod)
	fmt.Printf("%+v\n", mod)
	fmt.Println(res.StatusCode)
	fmt.Println(resStr)
}

func TestReqChmodVhostAndTopic(t *testing.T) {
	_, err := ReqCreateNewUser("user0", "user0")
	if err != nil {
		t.Error("An error occured but not this, instead:", err)
	}

	bodyStr := fmt.Sprintf(`{"exchange":"%s", "read":"^user0@a1f2f3b40c1d.*", "write":"^user0@a1f2f3b40c1d.*"}`, AMQP_TOPIC_EXCHANGE)
	body := bytes.NewBufferString(bodyStr)

	res, err := ReqChmodTopic(RABBITMQ_VHOST, "user0", body)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("Chmod Topic:", res.StatusCode)

	vhostMod := bytes.NewBufferString(`{"configure":".*","read":".*","write":".*"}`)

	res1, err := ReqChmodVhost(RABBITMQ_VHOST, "user0", vhostMod)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("Chmod Vhost:", res1.StatusCode)
}

func TestReqCreateNewQueueInVhost(t *testing.T) {
	sha := sha256.New()
	sha.Write([]byte(fmt.Sprint(time.Now().UTC().Unix())))
	devId := fmt.Sprintf("%x", sha.Sum(nil))

	// <username>@<dev-id>/air/temp
	// <username>@<dev-id>/air/humid
	topic0 := fmt.Sprintf("user0@%s/air/temp", devId)
	topic1 := fmt.Sprintf("user0@%s/air/humid", devId)

	res0, err := ReqCreateNewQueueInVhost(RABBITMQ_VHOST, topic0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ResponseReadString(res0))
	if res0.StatusCode != 201 {
		t.Error("Should be 201, got:", res0.StatusCode)
	}

	res1, err := ReqCreateNewQueueInVhost(RABBITMQ_VHOST, topic1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ResponseReadString(res1))
	if res1.StatusCode != 201 {
		t.Error("Should be 201, got:", res1.StatusCode)
	}
}

func TestReqCreateNewQueue(t *testing.T) {
	sha := sha256.New()
	sha.Write([]byte(fmt.Sprint(time.Now().UTC().Unix())))
	devId := fmt.Sprintf("%x", sha.Sum(nil))

	topic := fmt.Sprintf("user0@%s.air.temp", devId)

	res, err := ReqCreateNewQueue(topic)
	if err == nil && res.StatusCode == 201 {
		return
	}

	t.Error(err)
}

func TestReqBindExchangeRoutingToQueue(t *testing.T) {
	queue := "user0@dev1.air.humid"

	re1, err := ReqCreateNewQueue(queue)
	if err != nil {
		t.Error("An error occured but not this, instead:", err)
	}
	log.Println("QueueCreateStatus:", re1.StatusCode)
	log.Println("QueueCreateStatus:", ResponseReadString(re1))

	res, err := ReqBindExchangeRoutingToQueue(AMQP_TOPIC_EXCHANGE, queue, queue)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("status:", res.StatusCode)
	log.Println(ResponseReadString(res))
}

func TestReqBindAmqRoutingQueue(t *testing.T) {
	queue := "user0@dev1.air.humid"

	_, err := ReqCreateNewQueue(queue)
	if err != nil {
		t.Error("An error occured but not this, instead:", err)
	}

	res, err := ReqBindAmqRoutingQueue(queue, queue)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("status:", res.StatusCode)
	log.Println(ResponseReadString(res))
}

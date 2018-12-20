package main

import (
	"log"
	"fmt"
	"os"
	"github.com/streadway/amqp"
	"encoding/json"
	"encoding/base64"
	)

type Fax1 struct {
	FaxID string
	Caller int
	Wanted int
	File *os.File
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func main(){
	forever := make(chan bool)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	var info Fax1
	var way string
	var callernum =info.Caller
	var wantednum  =info.Wanted
	fax := make([]byte, 250)

	fmt.Println("Can you enter file document way")

	fmt.Scanf("%s", &way)
	fileway:=way

	fmt.Println("Can you enter caller number")

	fmt.Scanf("%d", &callernum)

	fmt.Println("Can you enter wanted number")

	fmt.Scanf("%d", &wantednum)

	fmt.Println(way)

	if fileway=="" {

		fmt.Println("Document way is empty!")

	}else {

	file, _ := os.Open(fileway)
	defer file.Close()
	readfile, _:=file.Read(fax)
	fmt.Println("readfile:",string(readfile))
	sEnc := map[string]string{"file": base64.StdEncoding.EncodeToString(fax)}
	fmt.Println("sEnc:",sEnc)
	fax_json, err := json.Marshal(sEnc)
	fmt.Println("fax_json:",fax_json)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body:  		 []byte (fax_json),
		})
	failOnError(err, "Failed to publish a message")
	<-forever
	}
}
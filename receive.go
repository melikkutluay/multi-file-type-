package main

import (
	"fmt"
	"log"
	"os"
	"github.com/streadway/amqp"
	"encoding/json"
	"os/exec"
	"encoding/base64"
)

type Fax struct{
	FileID int
	CallerNum int
	WantedNum int
	File []byte
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for compymsgs := range msgs {
			var fax Fax
			fmt.Println("dd:",compymsgs.Body)
			err := json.Unmarshal(compymsgs.Body, &fax)
			fmt.Println("",fax)
			if err != nil {
				fmt.Println("error:", err)
			}
			sendFile := string(fax.File)
			data, _ := base64.StdEncoding.DecodeString(sendFile)
			fmt.Println(data)
			//if err != nil {
			//	fmt.Println("error:", err)
			//}
			fmt.Println("\nSOLUTİONS\n")
			fmt.Println("s degeri:",sendFile,"\n")
			Create_fax_file(sendFile)
			exec.Command("echo")
			failOnError(err,"failed to execute command")
			log.Printf("Received a message: File:%s  ",fax.CallerNum,fax.WantedNum,fax.FileID,sendFile)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
func Create_fax_file(fax_text string)string {
	fireway:="C:/Users/melik/Desktop/ss/test_file.pdf"
	Faxer:=[]byte(fax_text)
	f,_:=os.Create(fireway)
	writer,_:=f.Write(Faxer)
	defer  f.Close()
	return string(writer)
}
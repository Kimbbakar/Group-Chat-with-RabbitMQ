package send

import (
	"bufio"
	"log"
	"os" 

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(username string) {
 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	  )
 
	scanner := bufio.NewScanner(os.Stdin)  
  
	log.Print("Enter the text: ")
	for {

		scanner.Scan()  
		input := scanner.Text() 
	
		body := input
		err = ch.Publish(
			"logs",     // exchange
			"", // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(username + ": " +  body),
			}) 
		failOnError(err, "Failed to publish a message")

	}


}
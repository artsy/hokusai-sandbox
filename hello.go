package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/streadway/amqp"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func ping (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func pub(rabbitmq_host string) {
	conn, err := amqp.Dial(rabbitmq_host)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hokusai-sandbox", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	for true {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
		time.Sleep(10 * time.Second)
	}
}

func sub(rabbitmq_host string) {
	conn, err := amqp.Dial(rabbitmq_host)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hokusai-sandbox", // name
		true,   // durable
		false,   // delete when unused
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
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/ping", ping)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	rabbitmq_host := os.Getenv("RABBITMQ_HOST")
	if len(rabbitmq_host) == 0 {
		rabbitmq_host = "amqp://guest:guest@localhost:5672/"
	}

	if os.Getenv("ENABLE_PUBLISH") != "" {
		go pub(rabbitmq_host)
	}

	if os.Getenv("ENABLE_SUBSCRIBE") != "" {
		go sub(rabbitmq_host)
	}

	fmt.Fprintf(os.Stderr, fmt.Sprintf("Listening on port %s...\n", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}

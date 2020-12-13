package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	product "projection/domain/product"
	mongoClient "projection/mongo"
)

func failOnError(err error, errMessage string) {
	if err != nil {
		log.Fatalf("%s: %s", errMessage, err)
	}
}

func main() {
	var (
		err 	error
		conn 	*amqp.Connection
		ch 		*amqp.Channel
		q 		 amqp.Queue
		msgs 	<-chan amqp.Delivery
	)

	log.Printf("RabbitMQ connection started")
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Printf("MongoDB connection started")
	err = mongoClient.LoadConfiguration()
	failOnError(err, "Failed to connect to MongoDB")

	log.Printf("Open channel service started")
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declare the queue from which we're going to consume
	q, err = ch.QueueDeclare(
		"InsertProduct",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err = ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			//convert payload to product model
			var productModel *product.Product
			if err = json.Unmarshal(d.Body, &productModel); err != nil {
				log.Printf("Payload can not converted to product model: %s", err)
				panic(err)
			}

			var repository product.ProductRepository = productModel
			if err = repository.InsertProduct(); err != nil {
				log.Printf("Record can not be inserted to db: %s", err)
				panic(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

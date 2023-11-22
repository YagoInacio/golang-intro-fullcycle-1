package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/infra/database"
	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/usecases"
	"github.com/yagoinacio/golang-intro-fullcycle-1/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)

	service := usecases.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgRabbitmqChannel) // Thread2 so it doesn't hold program's execution since it's a continuous process

	rabbitmqWorker(msgRabbitmqChannel, service)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, service *usecases.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecases.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}

		output, err := service.Execute(input)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Message received and stored on database:", output)
	}
}

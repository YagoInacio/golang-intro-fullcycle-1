package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/infra/database"
	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/usecases"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)

	service := usecases.NewCalculateFinalPrice(orderRepository)

	input := usecases.OrderInput{
		ID:    "2",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := service.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

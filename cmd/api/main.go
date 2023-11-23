package main

import (
	"net/http"

	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
	"github.com/yagoinacio/golang-intro-fullcycle-1/internal/entities"
)

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/orders", OrderHandler)
	// http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/orders", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderHandler(c echo.Context) error {
	order, _ := entities.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	order, _ := entities.NewOrder("1", 10, 1)
// 	err := order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	json.NewEncoder(w).Encode(order)
// }

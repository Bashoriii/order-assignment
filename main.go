package main

import (
	"fmt"
	"net/http"
	"order-assignment/database"
	"order-assignment/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	db := database.ConnectDatabase()
	defer db.Close()
	r := chi.NewRouter()

	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateOrder(w, r, db)
	})

	r.Get("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllOrders(w, r, db)
	})

	r.Get("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrderById(w, r, db, chi.URLParam(r, "id"))
	})

	r.Put("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateOrder(w, r, db, chi.URLParam(r, "id"))
	})

	r.Delete("/order/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteOrder(w, r, db, chi.URLParam(r, "id"))
	})

	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

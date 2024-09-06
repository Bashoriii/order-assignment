package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"order-assignment/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var order models.Order

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(order.CustomerName) > 30 {
		http.Error(w, "Customer name is too long", http.StatusBadRequest)
		return
	}

	// Validate each item in the order
	for i := range order.Items {
		// Input validation for name (max 25 characters)
		if len(order.Items[i].Name) > 30 {
			http.Error(w, "Item name is too long", http.StatusBadRequest)
			return
		}

		// Input validation for description (max 125 characters)
		if len(order.Items[i].Description) > 30 {
			http.Error(w, "Item description is too long", http.StatusBadRequest)
			return
		}
	}

	orderQuery := `INSERT INTO orders (customer_name) VALUES ($1) RETURNING id`
	err := db.QueryRow(orderQuery, order.CustomerName).Scan(&order.ID)
	if err != nil {
		http.Error(w, "Unable to create order", http.StatusInternalServerError)
		return
	}

	itemQuery := `INSERT INTO items (name, description, quantity, order_id) VALUES ($1, $2, $3, $4) RETURNING id`
	for i := range order.Items {
		err := db.QueryRow(itemQuery, order.Items[i].Name, order.Items[i].Description, order.Items[i].Quantity, order.ID).
			Scan(&order.Items[i].ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to create item: %v", err), http.StatusInternalServerError)
			return
		}
		order.Items[i].OrderId = order.ID
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

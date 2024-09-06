package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"order-assignment/models"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request, db *sql.DB, id string) {
	var updateOrder models.Order

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateOrderQuery := `UPDATE orders SET customer_name = $1 WHERE id = $2`
	_, err := db.Exec(updateOrderQuery, updateOrder.CustomerName, id)
	if err != nil {
		http.Error(w, "Unable to update order", http.StatusInternalServerError)
		return
	}

	// Insert new items into the items table
	itemQuery := `INSERT INTO items (name, description, quantity, order_id) VALUES ($1, $2, $3, $4) RETURNING id`
	for i := range updateOrder.Items {
		err := db.QueryRow(itemQuery, updateOrder.Items[i].Name, updateOrder.Items[i].Description, updateOrder.Items[i].Quantity, id).
			Scan(&updateOrder.Items[i].ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to add item: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Return the updated order as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateOrder)

}

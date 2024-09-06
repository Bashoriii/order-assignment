package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"order-assignment/models"
)

func GetOrderById(w http.ResponseWriter, r *http.Request, db *sql.DB, id string) {
	var order models.Order

	// Prepare the query
	orderQuery := `SELECT id, customer_name FROM orders WHERE id = $1`
	err := db.QueryRow(orderQuery, id).Scan(&order.ID, &order.CustomerName)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	itemQuery := `SELECT id, name, description, quantity FROM items WHERE order_id = $1`
	rows, err := db.Query(itemQuery, order.ID)
	if err != nil {
		http.Error(w, "Unable to retrieve items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		order.Items = append(order.Items, item)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode and send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"order-assignment/models"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	getAllQuery := `SELECT id, customer_name FROM orders`
	rows, err := db.Query(getAllQuery)
	if err != nil {
		http.Error(w, "Cant get all orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var order models.Order

		if err := rows.Scan(&order.ID, &order.CustomerName); err != nil {
			http.Error(w, "Unable to scan order", http.StatusInternalServerError)
			return
		}

		// Query to get items for the current order
		itemQuery := `SELECT id, name, description, quantity FROM items WHERE order_id = $1`
		itemRows, err := db.Query(itemQuery, order.ID)
		if err != nil {
			http.Error(w, "Unable to fetch items", http.StatusInternalServerError)
			return
		}
		defer itemRows.Close()

		// Collect items into a slice
		items := []models.Item{}
		for itemRows.Next() {
			var item models.Item
			err = itemRows.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity)
			if err != nil {
				http.Error(w, "Unable to scan item", http.StatusInternalServerError)
				return
			}
			item.OrderId = order.ID // Assign the order ID to the item
			items = append(items, item)
		}

		// Check for errors after iterating over rows
		if err = itemRows.Err(); err != nil {
			http.Error(w, "Error iterating over item rows", http.StatusInternalServerError)
			return
		}

		order.Items = items // Assign items to the order
		orders = append(orders, order)
	}

	// Check for errors after iterating over orders rows
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over order rows", http.StatusInternalServerError)
		return
	}

	// Encode the result as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

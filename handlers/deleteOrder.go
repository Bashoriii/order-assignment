package handlers

import (
	"database/sql"
	"net/http"
)

func DeleteOrder(w http.ResponseWriter, r *http.Request, db *sql.DB, id string) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	deleteItemsQuery := `DELETE FROM items WHERE order_id = $1`
	_, err := db.Exec(deleteItemsQuery, id)
	if err != nil {
		http.Error(w, "Unable to delete items related to the order", http.StatusInternalServerError)
		return
	}

	deleteOrderQuery := `DELETE FROM orders WHERE id = $1`
	result, err := db.Exec(deleteOrderQuery, id)
	if err != nil {
		http.Error(w, "Unable to delete order", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order and related items deleted successfully"))
}

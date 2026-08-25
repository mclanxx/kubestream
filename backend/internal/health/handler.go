package health

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

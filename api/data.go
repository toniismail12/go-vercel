package handler

import (
	"encoding/json"
	"fmt"
	"go-vercel/api/table"
	"net/http"
)

func Data(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var books []table.Books

	resp := map[string]any{
		"message": "data",
		"data":    books,
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}

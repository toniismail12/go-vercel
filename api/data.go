package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Books struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_url    string `json:"image_url"`
	Release_year string `json:"release_year"`
	Price        string `json:"price"`
	Total_pege   int    `json:"total_page"`
	Thickness    string `json:"thickness"`
	Created_at   time.Time
	Updated_at   time.Time
	Category_id  int `json:"category_id"`
}

func Data(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var books []Books

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

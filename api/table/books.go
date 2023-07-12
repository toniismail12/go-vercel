package table

import "time"

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

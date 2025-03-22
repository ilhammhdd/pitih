package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	type PaginatedResp[T any] struct {
		Page     int `json:"page"`
		PerPage  int `json:"perPage"`
		Total    int `json:"total"`
		Entities []T `json:"entities"`
	}
	type User struct {
		Email  string    `json:"email"`
		Joined time.Time `json:"joined"`
	}

	err := http.ListenAndServe(":8087", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header.Get("Authorization"))
		respBody := PaginatedResp[User]{
			Page:    1,
			PerPage: 10,
			Total:   3,
			Entities: []User{
				{Email: "user@example.com", Joined: time.Now()},
			},
			// Entities: []User{},
		}
		respBodyJSON, err := json.Marshal(respBody)
		if err != nil {
			log.Println(err)
		}
		w.Write(respBodyJSON)
	}))
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
}

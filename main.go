package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prettyApi/customDb"
	"prettyApi/customLog"
	"prettyApi/repository"
)

func main() {
	customLog.LogInit("./logs/app.log")
	customDb.Init()

	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		switch r.Method {
		case http.MethodGet:
			rep := repository.NewRepository()
			resp := rep.GetList()
			json.NewEncoder(w).Encode(resp)
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			var data map[string]interface{}
			err := decoder.Decode(&data)
			if err == nil {
				rep := repository.NewRepository()
				resp, err := rep.Create(data)
				if err != nil {
					json.NewEncoder(w).Encode(err.Error())
				} else {
					json.NewEncoder(w).Encode(resp)
				}
			} else {
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	})

	http.HandleFunc("/api/v1/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		switch r.Method {
		case http.MethodGet:
			rep := repository.NewRepository()
			id := r.PathValue("id")
			resp := rep.GetOne(id)
			json.NewEncoder(w).Encode(resp)
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

// dummy data
var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
	article{3, "data3", "content3"},
	article{4, "data4", "content4"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// mengubah format object ke json
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		// proses menambahkan article

		//proses penambahan berhasil
		responses := map[string]any{}
		responses["message"] = "data berhasil ditambahkan"
		responses["status"] = true
		var result, err = json.Marshal(responses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	// define endpoint
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/people", articles)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

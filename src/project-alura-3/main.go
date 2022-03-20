package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int32
}

var temp = template.Must(template.ParseGlob("Templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4200", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	product := []Product{
		{Name: "camisa", Description: "gola polo", Price: 19.90, Quantity: 30},
		{Name: "tenis", Description: "cano alto", Price: 199.90, Quantity: 10},
	}
	temp.ExecuteTemplate(w, "Index", product)
}

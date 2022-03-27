package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectWithDataBase() *sql.DB {
	conection := "user=postgres dbname=my_store password=1234567 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("Templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4200", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectWithDataBase()
	getAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err := getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.id = id
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()

}

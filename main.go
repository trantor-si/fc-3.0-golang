package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	name        string
	description string
	price       float64
}

func (p *Product) getSalesPrice() float64 {
	p.name = "trantor-si"
	return p.price * 2
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9191", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	product := Product{name: "iPhone", price: 1000}
	w.Write([]byte("Before insert..."))
	persistProduct(product)
	w.Write([]byte("Inserted!"))
}

func persistProduct(product Product) {
	db, err := sql.Open("sqlite3", "teste.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("Insert into products(name,purchasePrice) values ($1,$2)")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(product.name, product.price)
	if err != nil {
		fmt.Println(err)
	}
}

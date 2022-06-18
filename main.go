package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name        string  `json:"product-name" xml:"product-name"`
	Description string  `json:"product-description" xml:"product-description"`
	Price       float64 `json:"product-price" xml:"product-price"`
}

type User struct {
	Name  string `json:"user-name" xml:"user-name"`
	Email string `json:"user-email" xml:"user-email"`
}

var products []Product

func generateProducts() {
	p1 := Product{Name: "Product 1", Price: 1.99}
	p2 := Product{Name: "Product 1", Price: 2.99}
	products = append(products, p1, p2)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func listProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func listUsers(c echo.Context) error {
	u := &User{
		Name:  "Trantor SI",
		Email: "luccostajr@gmail.com",
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	generateProducts()

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users", listUsers)
	e.GET("/products", listProducts)

	e.Logger.Fatal(e.Start(":9191"))
}

// func persistProduct(product Product) {
// 	db, err := sql.Open("sqlite3", "teste.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	stmt, err := db.Prepare("Insert into products(name,price) values ($1,$2)")
// 	if err != nil {
// 		panic(err)
// 	}
// 	_, err = stmt.Exec(product.name, product.price)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

package main

import "fmt"

type Product struct {
	name          string
	description   string
	purchasePrice float64
}

func (p Product) getSalesPrice() float64 {
	p.name = "trantor-si"
	return p.purchasePrice * 2
}

func main() {
	product := Product{
		name:          "Laptop",
		description:   "Laptop description",
		purchasePrice: 1000.0,
	}

	fmt.Println(product.getSalesPrice(), ":", product.name)
}

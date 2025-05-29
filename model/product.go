package model

import (
	"fmt"
	"strings"
)

var ProductList []Product

type Product struct {
	Url, Image, Name, Price string
}

func (p *Product) PrintToScreen() {
	fmt.Println(strings.Repeat("=", 50), "PRODUCT")
	fmt.Printf("Product name	: %s\n", p.Name)
	fmt.Printf("Product price	: %s\n", p.Price)
	fmt.Printf("Product url		: %s\n", p.Url)
	fmt.Printf("Product image	: %s\n", p.Image)
	fmt.Println(strings.Repeat("=", 50))
}

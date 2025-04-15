package main

import "fmt"

// class
type Product struct {

	// attributes
	Id   int
	Name string
	Cost float64
}

// methods
func (p *Product) ToString() string {
	return fmt.Sprintf("Product :: Id = %d, Name = %q, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// Is-a relationship
type PerishableProduct struct {
	Product // (inheritance through composition)
	Expiry  string
}

// Overriding the ToString() of the Product
func (pp PerishableProduct) ToString() string {
	return fmt.Sprintf("%s, Expiry =%q", pp.Product.ToString(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		milk := PerishableProduct{
			Product: Product{
				Id:   101,
				Name: "Milk",
				Cost: 50,
			},
			Expiry: "2 Days",
		}
	*/
	milk := NewPerishableProduct(102, "Milk", 50, "2 Days")
	fmt.Printf("%#v\n", milk)
	fmt.Println(milk.Product.Id)
	fmt.Println(milk.Id) // accessing the attributes of the composed (inherited) type
	fmt.Println(milk.ToString())
}

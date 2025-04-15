/* standard library - https://pkg.go.dev/std */
package main

import "fmt"

// class
type Product struct {

	// attributes
	Id   int
	Name string
	Cost float64
}

// method
func (p *Product) ToString() string {
	return fmt.Sprintf("Product :: Id = %d, Name = %q, Cost = %0.02f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// Utility Function
func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

func main() {
	fmt.Println("Hello Go!")

	// create an object
	// p := Product{Id: 100, Name: "Pen", Cost: 10}
	pen := NewProduct(100, "Pen", 10)
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Println(pen.ToString())
	fmt.Println("After applying 10% discount")
	pen.ApplyDiscount(10)
	fmt.Println(pen.ToString())
}

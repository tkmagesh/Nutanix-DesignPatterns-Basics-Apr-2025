package main

import "fmt"

/* -------------------- */
/* Product Type */
/* -------------------- */
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

// Utility Function
func NewProduct(id int, name string, cost float64) *Product {
	return &Product{
		Id:   id,
		Name: name,
		Cost: cost,
	}
}

/* -------------------- */
/* Line Item */
/* -------------------- */

type LineItem struct {
	Product *Product //assciation
	Units   int
}

func NewLineItem(p *Product, units int) *LineItem {
	return &LineItem{
		Product: p,
		Units:   units,
	}
}

func (li LineItem) ToString() string {
	return fmt.Sprintf("%s, Units = %d", li.Product.ToString(), li.Units)
}

func (li LineItem) GetItemValue() float64 {
	return li.Product.Cost * float64(li.Units)
}

/* -------------------- */
/* Shopping Cart Type */
/* -------------------- */
type ShoppingCart struct {
	LineItems []*LineItem // composition
}

func (sc *ShoppingCart) AddItem(p *Product, units int) {
	newLineItem := NewLineItem(p, units)
	sc.LineItems = append(sc.LineItems, newLineItem)
}

func (sc ShoppingCart) GetCartValue() float64 {
	var cartValue float64
	for _, li := range sc.LineItems {
		cartValue += li.GetItemValue()
	}
	return cartValue
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		LineItems: make([]*LineItem, 0),
	}
}

func main() {
	sc := NewShoppingCart()
	sc.AddItem(NewProduct(101, "Pen", 10), 10)
	sc.AddItem(NewProduct(102, "Pencil", 5), 20)
	sc.AddItem(NewProduct(103, "Marker", 50), 5)

	for _, li := range sc.LineItems {
		fmt.Println(li.ToString())
	}
	fmt.Printf("Cart Value : %0.02f\n", sc.GetCartValue())
}

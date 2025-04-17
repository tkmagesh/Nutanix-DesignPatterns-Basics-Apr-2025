package main

import "fmt"

// Product
type Pizza struct {
	dough    string
	sauce    string
	toppings []string
}

// Builder Interface
type PizzaBuilder interface {
	SetDough(string) PizzaBuilder
	SetSauce(string) PizzaBuilder
	AddTopping(string) PizzaBuilder
	Build() Pizza
}

// Concrete Builder
type CustomPizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() *CustomPizzaBuilder {
	return &CustomPizzaBuilder{pizza: Pizza{}}
}

func (b *CustomPizzaBuilder) SetDough(dough string) PizzaBuilder {
	b.pizza.dough = dough
	return b
}

func (b *CustomPizzaBuilder) SetSauce(sauce string) PizzaBuilder {
	b.pizza.sauce = sauce
	return b
}

func (b *CustomPizzaBuilder) AddTopping(topping string) PizzaBuilder {
	b.pizza.toppings = append(b.pizza.toppings, topping)
	return b
}

func (b *CustomPizzaBuilder) Build() Pizza {
	return b.pizza
}

// Client
func main() {

	pizza := NewPizzaBuilder().
		SetDough("Thin Crust").
		SetSauce("Tomato").
		AddTopping("Cheese").
		AddTopping("Olives").
		Build()

	fmt.Printf("Pizza: %+v\n", pizza)
}

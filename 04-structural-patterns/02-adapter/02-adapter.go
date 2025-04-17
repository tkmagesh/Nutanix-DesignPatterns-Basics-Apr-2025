package main

import "fmt"

// Target interface
type PaymentProcessor interface {
	Pay(amount float64)
}

// Adaptees
type PayPalService struct{}

func (p *PayPalService) SendPayment(amount float64) {
	fmt.Printf("PayPal processed payment of %.2f\n", amount)
}

type StripeService struct{}

func (s *StripeService) ChargeCard(amount float64) {
	fmt.Printf("Stripe charged %.2f\n", amount)
}

type RazorpayService struct{}

func (r *RazorpayService) MakeTransaction(amount float64) {
	fmt.Printf("Razorpay transaction of %.2f\n", amount)
}

// Adapters
type PayPalAdapter struct{ service *PayPalService }

func (a *PayPalAdapter) Pay(amount float64) {
	a.service.SendPayment(amount)
}

type StripeAdapter struct{ service *StripeService }

func (a *StripeAdapter) Pay(amount float64) {
	a.service.ChargeCard(amount)
}

type RazorpayAdapter struct{ service *RazorpayService }

func (a *RazorpayAdapter) Pay(amount float64) {
	a.service.MakeTransaction(amount)
}

// Client
func processPayment(p PaymentProcessor, amount float64) {
	p.Pay(amount)
}

func main() {
	payments := []PaymentProcessor{
		&PayPalAdapter{&PayPalService{}},
		&StripeAdapter{&StripeService{}},
		&RazorpayAdapter{&RazorpayService{}},
	}

	for _, p := range payments {
		processPayment(p, 1000)
	}
}

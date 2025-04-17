package main

import "fmt"

// Common interface
type PaymentProcessor interface {
	Pay(amount float64)
}

// Adaptees
type PayPalService struct{}

func (p *PayPalService) SendPayment(amount float64) {
	fmt.Printf("PayPal: %.2f paid\n", amount)
}

type StripeService struct{}

func (s *StripeService) ChargeCard(amount float64) {
	fmt.Printf("Stripe: %.2f charged\n", amount)
}

type RazorpayService struct{}

func (r *RazorpayService) MakeTransaction(amount float64) {
	fmt.Printf("Razorpay: %.2f transacted\n", amount)
}

// Adapters
type PayPalAdapter struct{ svc *PayPalService }

func (a *PayPalAdapter) Pay(amount float64) {
	a.svc.SendPayment(amount)
}

type StripeAdapter struct{ svc *StripeService }

func (a *StripeAdapter) Pay(amount float64) {
	a.svc.ChargeCard(amount)
}

type RazorpayAdapter struct{ svc *RazorpayService }

func (a *RazorpayAdapter) Pay(amount float64) {
	a.svc.MakeTransaction(amount)
}

// Factory
func GetPaymentProcessor(gateway string) (PaymentProcessor, error) {
	switch gateway {
	case "paypal":
		return &PayPalAdapter{&PayPalService{}}, nil
	case "stripe":
		return &StripeAdapter{&StripeService{}}, nil
	case "razorpay":
		return &RazorpayAdapter{&RazorpayService{}}, nil
	default:
		return nil, fmt.Errorf("unsupported payment gateway: %s", gateway)
	}
}

// Client
func Checkout(amount float64, gateway string) {
	processor, err := GetPaymentProcessor(gateway)
	if err != nil {
		fmt.Println(err)
		return
	}
	processor.Pay(amount)
}

func main() {
	Checkout(1000, "paypal")
	Checkout(2000, "stripe")
	Checkout(3000, "razorpay")
}

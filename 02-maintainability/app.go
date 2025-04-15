package main

import (
	"fmt"
	"time"
)

type TimeProvider struct {
}

func (tp TimeProvider) GetCurrentTime() time.Time {
	return time.Now()
}

func NewTimeProvider() *TimeProvider {
	return &TimeProvider{}
}

/*
Responsibility
- decides what message to show
*/

type ITimeProvider interface {
	GetCurrentTime() time.Time
}

type Greeter struct {
	UserName string
	tp       ITimeProvider
}

func (greeter Greeter) Greet() string {
	// get the current time
	// currentTime := time.Now()

	// tp := NewTimeProvider() //violation of DIP

	currentTime := greeter.tp.GetCurrentTime()

	// display a greet message depending on the time
	if currentTime.Hour() < 12 {
		return fmt.Sprintf("Hi %s, Good Morning!\n", greeter.UserName)
	} else {
		return fmt.Sprintf("Hi %s, Good Day!\n", greeter.UserName)
	}
}

func NewGreeter(tp ITimeProvider) *Greeter {
	return &Greeter{
		tp: tp,
	}
}

func main() {
	/*
		Responsibilities
			- interfacing with the user (accepting inputs, displaying outputs)

	*/
	var userName string

	// accept the userName
	fmt.Println("Enter your name :")
	fmt.Scanln(&userName)

	tp := NewTimeProvider()
	greeter := NewGreeter(tp)

	greeter.UserName = userName
	msg := greeter.Greet()
	fmt.Println(msg)

}

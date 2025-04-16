package main

import "fmt"

// Abstract Products
type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

// Concrete Products - Windows
type WindowsButton struct{}

func (WindowsButton) Paint() { fmt.Println("Rendering Windows Button") }

type WindowsCheckbox struct{}

func (WindowsCheckbox) Paint() { fmt.Println("Rendering Windows Checkbox") }

// Concrete Products - MacOS
type MacButton struct{}

func (MacButton) Paint() { fmt.Println("Rendering Mac Button") }

type MacCheckbox struct{}

func (MacCheckbox) Paint() { fmt.Println("Rendering Mac Checkbox") }

// Abstract Factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factories
type WindowsFactory struct{} //implements GUIFactory

func (WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}
func (WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

type MacFactory struct{} //implements GUIFactory

func (MacFactory) CreateButton() Button {
	return MacButton{}
}
func (MacFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}

// Client
func renderUI(factory GUIFactory) {
	btn := factory.CreateButton()
	chk := factory.CreateCheckbox()
	btn.Paint()
	chk.Paint()
}

func main() {
	var factory GUIFactory

	// simulate OS
	os := "mac"
	if os == "windows" {
		factory = WindowsFactory{}
	} else {
		factory = MacFactory{}
	}

	renderUI(factory)
}

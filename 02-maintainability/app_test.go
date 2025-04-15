package main

import (
	"testing"
	"time"
)

// simulate time provider return morning time
type MorningTimeProvider struct {
}

func (mtp MorningTimeProvider) GetCurrentTime() time.Time {
	return time.Date(2025, time.April, 15, 9, 0, 0, 0, time.UTC)
}

// simulate time provider return morning time
type AfternoonTimeProvider struct {
}

func (atp AfternoonTimeProvider) GetCurrentTime() time.Time {
	return time.Date(2025, time.April, 15, 15, 0, 0, 0, time.UTC)
}

func Test_Greeter_Before_12(t *testing.T) {
	// arrange
	userName := "Magesh"
	tp := MorningTimeProvider{}
	greeter := NewGreeter(tp)
	greeter.UserName = userName
	expected := "Hi Magesh, Good Morning!\n"

	// act
	actual := greeter.Greet()

	// assert
	if expected != actual {
		t.Errorf("Expected : %q, Actual : %q", expected, actual)
	}
}

func Test_Greeter_After_12(t *testing.T) {
	// arrange
	userName := "Magesh"
	tp := AfternoonTimeProvider{}
	greeter := NewGreeter(tp)
	greeter.UserName = userName
	expected := "Hi Magesh, Good Day!\n"

	// act
	actual := greeter.Greet()

	// assert
	if expected != actual {
		t.Errorf("Expected : %q, Actual : %q", expected, actual)
	}
}

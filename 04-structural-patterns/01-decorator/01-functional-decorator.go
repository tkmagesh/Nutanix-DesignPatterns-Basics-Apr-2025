package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	logAdd := logDecorator(add)
	logSubtract := logDecorator(subtract)

	profileAdd := profileDecorator(logAdd)
	profileSubtract := profileDecorator(logSubtract)
	profileAdd(100, 200)
	profileSubtract(100, 200)
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// v2.0
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

// applying commonality-variability on the above
func logOperation(operation func(int, int), x, y int) {
	log.Println("Operation started")
	operation(x, y)
	log.Println("Operation completed")
}

// decorator using function composition

func logDecorator(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		operation(x, y)
		log.Println("Operation completed")
	}
}

func profileDecorator(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

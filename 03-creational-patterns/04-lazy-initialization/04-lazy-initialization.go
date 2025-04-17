package main

import (
	"fmt"
	"sync"
)

type LazyResource struct {
	data []int
	once sync.Once
}

func (lr *LazyResource) GetData() []int {
	lr.once.Do(func() {
		fmt.Println("Initializing data...")
		lr.data = make([]int, 1000000)
		for i := range lr.data {
			lr.data[i] = i * i
		}
	})
	return lr.data
}

func main() {
	resource := &LazyResource{}
	fmt.Println("Before accessing data")
	fmt.Println(resource.GetData()[:5]) // Data gets initialized here
	fmt.Println("GetData() again....")
	fmt.Println(resource.GetData()[:5]) // Data gets initialized here
}

package main

import (
	"fmt"
	"sync"
)

type RegionalDBConnection struct {
	region string
}

var (
	multiton = make(map[string]*RegionalDBConnection)
	mutex    sync.Mutex
)

func GetDBConnection(region string) *RegionalDBConnection {
	mutex.Lock()
	defer mutex.Unlock()

	if conn, ok := multiton[region]; ok {
		return conn
	}
	conn := &RegionalDBConnection{region: region}
	multiton[region] = conn
	return conn
}

func main() {
	us := GetDBConnection("US")
	eu := GetDBConnection("EU")
	anotherUS := GetDBConnection("US")

	fmt.Println(us == anotherUS) // true
}

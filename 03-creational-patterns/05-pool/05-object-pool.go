package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	ID int
}

func (c *Connection) Connect() {
	fmt.Printf("Connection %d in use.\n", c.ID)
}

type ConnectionPool struct {
	mu        sync.Mutex
	available []*Connection
	inUse     map[int]*Connection
}

func NewConnectionPool(size int) *ConnectionPool {
	pool := &ConnectionPool{
		available: make([]*Connection, size),
		inUse:     make(map[int]*Connection),
	}
	for i := 0; i < size; i++ {
		pool.available[i] = &Connection{ID: i}
	}
	return pool
}

func (p *ConnectionPool) Acquire() (*Connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.available) == 0 {
		return nil, fmt.Errorf("no available connections")
	}
	conn := p.available[0]
	p.available = p.available[1:]
	p.inUse[conn.ID] = conn
	return conn, nil
}

func (p *ConnectionPool) Release(conn *Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.inUse, conn.ID)
	p.available = append(p.available, conn)
}

func main() {
	pool := NewConnectionPool(2)

	conn1, _ := pool.Acquire()
	conn1.Connect()

	pool.Release(conn1)
}

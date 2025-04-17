package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

// Logger
type Logger struct{}

func (l *Logger) Log(msg string) {
	fmt.Printf("[LOG]: %s\n", msg)
}

// Cache
type Cache struct {
	mu       sync.RWMutex
	products []Product
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (c *Cache) Set(data []Product) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products = data
	fmt.Println("[CACHE]: Updated")
}

func (c *Cache) Get() []Product {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println("[CACHE]: Retrieved")
	return c.products
}

func (c *Cache) Invalidate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products = nil
	fmt.Println("[CACHE]: Invalidated")
}

// DB Connection
type DBConnection struct {
	logger *Logger
}

func (db *DBConnection) Connect() *sql.DB {
	db.logger.Log("Connecting to DB")
	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// Product Repository
type ProductRepository struct {
	conn   *sql.DB
	logger *Logger
}

func NewProductRepository(conn *sql.DB, logger *Logger) *ProductRepository {
	_, _ = conn.Exec(`CREATE TABLE products (id INTEGER, name TEXT, price REAL)`)
	logger.Log("Created table products")
	return &ProductRepository{conn: conn, logger: logger}
}

func (r *ProductRepository) AddProduct(id int, name string, price float64) {
	r.logger.Log(fmt.Sprintf("Adding product: %s", name))
	_, _ = r.conn.Exec(`INSERT INTO products VALUES (?, ?, ?)`, id, name, price)
}

func (r *ProductRepository) ListProducts() []Product {
	r.logger.Log("Fetching products from DB")
	rows, _ := r.conn.Query(`SELECT id, name, price FROM products`)
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		_ = rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	return products
}

// Product Service (Facade)
type ProductService struct {
	db     *sql.DB
	logger *Logger
	repo   *ProductRepository
	cache  *Cache
}

func NewProductService() *ProductService {
	logger := &Logger{}
	conn := (&DBConnection{logger: logger}).Connect()
	repo := NewProductRepository(conn, logger)
	cache := &Cache{}
	return &ProductService{
		db: conn, logger: logger, repo: repo, cache: cache,
	}
}

func (s *ProductService) AddProduct(id int, name string, price float64) {
	s.logger.Log("Validating product...")
	s.repo.AddProduct(id, name, price)
	s.cache.Invalidate()
	s.logger.Log("Product added successfully")
}

func (s *ProductService) ListAll() {
	cached := s.cache.Get()
	if cached != nil {
		for _, p := range cached {
			fmt.Printf("From Cache - ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
		}
		return
	}
	products := s.repo.ListProducts()
	s.cache.Set(products)
	for _, p := range products {
		fmt.Printf("From DB - ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

// Client
func main() {
	service := NewProductService()
	service.AddProduct(1, "Laptop", 1200)
	service.AddProduct(2, "Mouse", 25)

	service.ListAll() // From DB
	service.ListAll() // From Cache
}

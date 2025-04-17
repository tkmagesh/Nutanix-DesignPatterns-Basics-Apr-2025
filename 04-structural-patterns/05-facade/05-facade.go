package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// Subsystem: DB Connection
type DBConnection struct{}

func (db *DBConnection) Connect() *sql.DB {
	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// Subsystem: Product Operations
type ProductRepository struct {
	conn *sql.DB
}

func NewProductRepository(conn *sql.DB) *ProductRepository {
	_, _ = conn.Exec(`CREATE TABLE products (id INTEGER, name TEXT, price REAL)`)
	return &ProductRepository{conn: conn}
}

func (r *ProductRepository) AddProduct(id int, name string, price float64) {
	_, _ = r.conn.Exec(`INSERT INTO products VALUES (?, ?, ?)`, id, name, price)
}

func (r *ProductRepository) ListProducts() {
	rows, _ := r.conn.Query(`SELECT id, name, price FROM products`)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var price float64
		_ = rows.Scan(&id, &name, &price)
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", id, name, price)
	}
}

// Facade
type ProductService struct {
	db   *sql.DB
	repo *ProductRepository
}

func NewProductService() *ProductService {
	conn := (&DBConnection{}).Connect()
	repo := NewProductRepository(conn)
	return &ProductService{db: conn, repo: repo}
}

func (s *ProductService) AddProduct(id int, name string, price float64) {
	fmt.Println("Validating product...")
	s.repo.AddProduct(id, name, price)
	fmt.Println("Product added.")
}

func (s *ProductService) ListAll() {
	s.repo.ListProducts()
}

// Client
func main() {
	service := NewProductService()
	service.AddProduct(1, "Laptop", 1200)
	service.AddProduct(2, "Mouse", 25)

	service.ListAll()
}


import sqlite3

# Subsystem: DB Connection
class DBConnection:
    def connect(self):
        return sqlite3.connect(":memory:")

# Subsystem: Product Operations
class ProductRepository:
    def __init__(self, conn):
        self.conn = conn
        self.create_table()

    def create_table(self):
        self.conn.execute("CREATE TABLE products (id INTEGER, name TEXT, price REAL)")

    def add_product(self, id, name, price):
        self.conn.execute("INSERT INTO products VALUES (?, ?, ?)", (id, name, price))
        self.conn.commit()

    def list_products(self):
        return self.conn.execute("SELECT * FROM products").fetchall()

# Facade
class ProductService:
    def __init__(self):
        self.db = DBConnection()
        self.conn = self.db.connect()
        self.repo = ProductRepository(self.conn)

    def add_product(self, id, name, price):
        print("Validating product...")
        self.repo.add_product(id, name, price)
        print("Product added.")

    def list_all(self):
        return self.repo.list_products()

# Client
service = ProductService()
service.add_product(1, "Laptop", 1200)
service.add_product(2, "Mouse", 25)

print(service.list_all())
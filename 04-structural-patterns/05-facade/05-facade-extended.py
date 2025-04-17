import sqlite3
from functools import lru_cache

class Logger:
    @staticmethod
    def log(message):
        print(f"[LOG]: {message}")

class DBConnection:
    def connect(self):
        Logger.log("Connecting to DB")
        return sqlite3.connect(":memory:")

class ProductRepository:
    def __init__(self, conn):
        self.conn = conn
        self.create_table()

    def create_table(self):
        Logger.log("Creating products table")
        self.conn.execute("CREATE TABLE products (id INTEGER, name TEXT, price REAL)")

    def add_product(self, id, name, price):
        Logger.log(f"Adding product: {name}")
        self.conn.execute("INSERT INTO products VALUES (?, ?, ?)", (id, name, price))
        self.conn.commit()

    def list_products(self):
        Logger.log("Fetching products from DB")
        return self.conn.execute("SELECT * FROM products").fetchall()

class Cache:
    def __init__(self):
        self.products = None

    def set(self, data):
        Logger.log("Updating cache")
        self.products = data

    def get(self):
        Logger.log("Retrieving from cache")
        return self.products

    def invalidate(self):
        Logger.log("Invalidating cache")
        self.products = None

class ProductService:
    def __init__(self):
        self.db = DBConnection()
        self.conn = self.db.connect()
        self.repo = ProductRepository(self.conn)
        self.cache = Cache()

    def add_product(self, id, name, price):
        Logger.log("Validating product...")
        self.repo.add_product(id, name, price)
        self.cache.invalidate()
        Logger.log("Product added successfully")

    def list_all(self):
        cached = self.cache.get()
        if cached is not None:
            return cached
        products = self.repo.list_products()
        self.cache.set(products)
        return products

# Client
service = ProductService()
service.add_product(1, "Laptop", 1200)
service.add_product(2, "Mouse", 25)
print(service.list_all())   # From DB
print(service.list_all())   # From Cache

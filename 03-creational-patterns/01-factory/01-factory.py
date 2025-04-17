from abc import ABC, abstractmethod

# Product interface
class DBConnection(ABC):
    @abstractmethod
    def connect(self):
        pass

# Concrete products
class MySQLConnection(DBConnection):
    def connect(self):
        return "Connected to MySQL"

class PostgresConnection(DBConnection):
    def connect(self):
        return "Connected to PostgreSQL"

# Factory class
class DBFactory:
    def create_connection(self, db_type: str) -> DBConnection:
        if db_type == "mysql":
            return MySQLConnection()
        elif db_type == "postgres":
            return PostgresConnection()
        else:
            raise ValueError("Unknown DB type")

# Usage
factory = DBFactory()

conn1 = factory.create_connection("mysql")
print(conn1.connect())

conn2 = factory.create_connection("postgres")
print(conn2.connect())

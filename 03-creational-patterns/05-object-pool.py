import threading

class Connection:
    def __init__(self, id):
        self.id = id

    def connect(self):
        print(f"Connection {self.id} in use.")

class ConnectionPool:
    def __init__(self, size=3):
        self.lock = threading.Lock()
        self.available = [Connection(i) for i in range(size)]
        self.in_use = []

    def acquire(self):
        with self.lock:
            if not self.available:
                raise Exception("No available connections")
            conn = self.available.pop()
            self.in_use.append(conn)
            return conn

    def release(self, conn):
        with self.lock:
            self.in_use.remove(conn)
            self.available.append(conn)

# Usage
pool = ConnectionPool()

conn1 = pool.acquire()
conn1.connect()

pool.release(conn1)
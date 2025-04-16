###  Object Pool Pattern 

The **Object Pool Pattern** is a **creational design pattern** used to manage the reuse of **expensive-to-create objects**. Instead of creating and destroying objects repeatedly, a pool (like a cache) holds a set of initialized objects ready to use, improving performance and resource usage.

---

###  When to Use It
- Object creation is **costly** (e.g., database connections, threads, sockets).
- You need **many short-lived objects**.
- System is **resource constrained**.

---

###  Key Concepts
- **Pool (Manager):** Maintains the list of available and in-use objects.
- **Reusable Object:** The object managed by the pool.
- **Acquire:** Get an object from the pool.
- **Release:** Return the object to the pool.
# Proxy Pattern
The **Proxy Pattern** is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. It is commonly used to add additional behavior (like access control, lazy loading, logging, or remote access) without changing the original object's code.

---

### **Real-World Example Use Case**
- **Virtual Proxy**: Delay the creation and loading of a heavy object until it's needed.
- **Remote Proxy**: Represent an object in a different address space (e.g., accessing a remote service).
- **Protection Proxy**: Add access control to sensitive objects.


## Proxy Vs Adapter
Here's a concise comparison of the **Proxy** and **Adapter** design patterns:

| Aspect              | **Proxy**                                   | **Adapter**                                     |
|---------------------|---------------------------------------------|-------------------------------------------------|
| **Intent**           | Control access to an object                 | Convert one interface to another                |
| **Focus**            | Adds additional behavior (e.g., access control, lazy init) | Interface compatibility                        |
| **Client Interface** | Same as the real object                     | Different from the adaptee                      |
| **Used For**         | Access control, logging, lazy instantiation, remote access | Integration with incompatible interfaces        |
| **Example**          | Virtual proxy for image loading             | Convert legacy API to match modern interface    |
| **Pattern Type**     | Structural                                   | Structural                                      |

---

### Real-World Analogy

- **Proxy**: Like a receptionist who checks if you can meet the CEO.
- **Adapter**: Like a travel plug adapter that lets a US plug fit in an EU socket.





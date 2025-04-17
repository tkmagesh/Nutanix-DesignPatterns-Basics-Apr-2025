# Mediator Pattern

The **Mediator Pattern** is a behavioral design pattern that **centralizes communication between objects** to reduce dependencies between them. Instead of objects referring to each other directly, they communicate via a **mediator object**, which promotes loose coupling and improves code maintainability.

---

### **Intent**
> Define an object that encapsulates how a set of objects interact. Promotes loose coupling by keeping objects from referring to each other explicitly.

---

### **Real-World Analogy**
Think of an **air traffic control system**. Planes (objects) don't talk to each other directly. Instead, they communicate through the control tower (mediator), which coordinates landing and takeoff to avoid collisions.



### **Use Cases**
- Chatroom where participants send messages via a central server
- UI components interacting via a dialog controller
- Event handling systems (like JavaScript DOM events)


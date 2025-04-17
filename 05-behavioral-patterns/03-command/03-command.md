# Command Pattern

The **Command Pattern** is a behavioral design pattern used to encapsulate a request as an object, thereby allowing you to parameterize clients with queues, requests, and operations. It's particularly useful when you need to issue requests to objects without knowing anything about the operation being requested or the receiver of the request.

### **Intent**
- Encapsulate all the details of a command in a single object.
- Decouple the object that invokes the operation from the one that knows how to perform it.
- Allow support for undoable operations, queuing commands, and logging history.

---

### **Structure**

- **Command** – declares an interface for executing an operation.
- **ConcreteCommand** – defines a binding between a Receiver and an action.
- **Receiver** – knows how to perform the work needed to carry out a request.
- **Invoker** – asks the command to carry out the request.
- **Client** – creates a command object and sets its receiver.

---

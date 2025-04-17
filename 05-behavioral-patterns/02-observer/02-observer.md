The **Observer Pattern** is a behavioral design pattern where an object, called the **subject**, maintains a list of its dependents, called **observers**, and notifies them automatically of any state changes.

This pattern is commonly used when changes to one object need to be propagated to others without tight coupling between them.

---

### **Real-world Analogy**
A news agency (**Subject**) notifies its subscribers (**Observers**) whenever a new article is published. Subscribers can subscribe/unsubscribe independently.

---

### **Structure**
- **Subject**: Holds a list of observers and provides methods to attach, detach, and notify them.
- **Observer**: Defines an interface for objects that should be notified of changes.
- **ConcreteSubject**: Stores state and notifies observers on change.
- **ConcreteObserver**: Implements the update interface to react to changes.

---


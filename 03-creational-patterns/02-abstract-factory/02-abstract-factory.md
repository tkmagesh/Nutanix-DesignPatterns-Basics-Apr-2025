# Abstract Factory

The **Abstract Factory** pattern is a **creational design pattern** that provides an interface for creating families of related or dependent objects **without specifying their concrete classes**.

###  Key Idea
Instead of instantiating classes directly, you use **factories** that produce objects. When you want to switch the "family" (e.g., UI components for Mac or Windows, or connectors for MySQL or PostgreSQL), you just change the factory — not the product-using code.

---

###  Structure

1. **AbstractFactory**  
   Defines methods to create abstract products.

2. **ConcreteFactory**  
   Implements the methods of the AbstractFactory to create concrete product objects.

3. **AbstractProduct**  
   Declares interfaces for a set of products.

4. **ConcreteProduct**  
   Implements the AbstractProduct interface.

5. **Client**  
   Uses only interfaces (not concrete classes) to work with products.

---

###  Example – GUI Toolkit

| Component        | Example                        |
|------------------|--------------------------------|
| AbstractFactory  | `GUIFactory`                   |
| ConcreteFactory  | `MacFactory`, `WindowsFactory` |
| AbstractProduct  | `Button`, `Checkbox`           |
| ConcreteProduct  | `MacButton`, `WinCheckbox`     |
| Client           | `renderUI(factory)`            |

---

###  When to Use

- When you want to create families of related objects (e.g., different themes or OS-specific components).
- When system configuration needs to switch object families at runtime.
- When enforcing consistency between related objects (e.g., same "theme" or "style") is necessary.


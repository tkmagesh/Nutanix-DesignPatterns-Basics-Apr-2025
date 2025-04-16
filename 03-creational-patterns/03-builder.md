### Builder Pattern

The **Builder Pattern** is a **creational design pattern** that **separates the construction of a complex object from its representation**, allowing the same construction process to create different representations.

---

### Key Features

- **Step-by-step construction**: Useful when the object requires numerous steps or options to construct.
- **Immutable result**: Builder often produces immutable objects.
- **Readable DSL-style code**: Enables fluent interfaces like `builder.setA().setB().build()`.

---

### Structure

1. **Product**: The complex object being built.
2. **Builder**: Interface declaring building steps.
3. **ConcreteBuilder**: Implements the steps to construct the object.
4. **Director** *(optional)*: Controls the order of steps.

---

### When to Use

- When creating complex objects with many optional parts.
- When object creation should be independent of its parts.
- To avoid constructors with too many parameters (**Telescoping Constructor Anti-pattern**).

---

### Real-World Analogies

- Ordering a pizza with optional toppings (like Domino’s or Zomato UI).
- Building a query with optional clauses (e.g., SQL Query Builder).
- Creating HTML documents using tag-by-tag assembly.


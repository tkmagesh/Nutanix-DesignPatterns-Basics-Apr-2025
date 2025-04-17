# Chain of Responsibility

The **Chain of Responsibility** is a **behavioral design pattern** used to pass a request along a chain of handlers. Each handler decides whether to process the request or pass it to the next handler in the chain. This decouples the sender of the request from its potential receivers.

### Key Concepts:
- **Sender** doesn't know which object will handle the request.
- **Handlers** are linked, usually in a chain or pipeline.
- Each handler either handles the request or delegates it to the next.

---

### Real-World Analogy:
Consider **technical support** at a company:
- Tier 1 handles basic issues.
- If it can't resolve the problem, it escalates to Tier 2.
- Tier 2 may escalate to Tier 3, and so on.

---

### Structure:
```text
Client → Handler1 → Handler2 → Handler3 → ...
```

---

### Example Use Cases:
- Logging (different log levels).
- Event propagation in GUI frameworks.
- Access control mechanisms (e.g., middleware in web frameworks).

## Difference between Chain Of Responsibility & Decorator


### 1. Intent

| Aspect                 | Chain of Responsibility                     | Decorator                               |
|------------------------|---------------------------------------------|------------------------------------------|
| **Purpose**            | Pass a request along a chain of handlers    | Add or modify behavior of an object      |
| **Responsibility**     | One or more handlers may handle the request | Each decorator adds new behavior         |
| **Main goal**          | Flexibility in who handles a request        | Dynamic behavior composition             |

---

### 2. Structure

| Aspect                 | Chain of Responsibility                     | Decorator                               |
|------------------------|---------------------------------------------|------------------------------------------|
| **Flow Direction**     | Linear passing until handled                | Nested wrapping, all decorators invoked  |
| **Handlers / Wrappers**| Each handler has a reference to the next    | Each wrapper has a reference to the component |
| **Termination**        | May stop at first handler that processes it | Always flows through all wrappers        |

---

### 3. Behavior

| Aspect                 | Chain of Responsibility                     | Decorator                               |
|------------------------|---------------------------------------------|------------------------------------------|
| **Control Flow**       | Stops once a handler handles the request    | All decorators are always executed       |
| **Typical Use**        | Event handling, middleware, command routing | Adding logging, caching, compression     |

---


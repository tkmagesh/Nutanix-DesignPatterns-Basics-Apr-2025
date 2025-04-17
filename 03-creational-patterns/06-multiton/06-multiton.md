The **Multiton Pattern** is a creational design pattern that generalizes the Singleton pattern. While Singleton ensures only *one* instance exists for a class, **Multiton ensures only *one instance per key*** exists.

###  When to Use Multiton
- You need to manage a fixed set of instances (like per database, user session, language).
- You want to avoid creating the same object repeatedly for the same key.

---

###  Core Idea

> Instead of creating a new instance every time, store instances in a map keyed by an identifier, and return the same instance for the same key.

---

###  Real-World Example

#### Scenario: DB connections per region

- You want one DB connection per region: `"US"`, `"EU"`, `"ASIA"`.
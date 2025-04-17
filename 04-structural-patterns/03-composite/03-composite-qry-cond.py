from abc import ABC, abstractmethod

class Expression(ABC):
    @abstractmethod
    def to_sql(self):
        pass

class Condition(Expression):
    def __init__(self, field, operator, value):
        self.field = field
        self.operator = operator
        self.value = value

    def to_sql(self):
        return f"{self.field} {self.operator} '{self.value}'"

class AndExpression(Expression):
    def __init__(self, *children):
        self.children = children

    def to_sql(self):
        return "(" + " AND ".join(child.to_sql() for child in self.children) + ")"

class OrExpression(Expression):
    def __init__(self, *children):
        self.children = children

    def to_sql(self):
        return "(" + " OR ".join(child.to_sql() for child in self.children) + ")"

if __name__ == "__main__":
    cond1 = Condition("age", ">", "30")
    cond2 = Condition("city", "=", "Delhi")
    cond3 = Condition("salary", ">", "50000")

    and_expr = AndExpression(cond1, cond2)
    or_expr = OrExpression(and_expr, cond3)

    print("WHERE", or_expr.to_sql())

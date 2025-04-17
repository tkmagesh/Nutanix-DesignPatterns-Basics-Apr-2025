""" Convert the query builder to adapt fluent interface """

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

class And(Expression):
    def __init__(self, *children):
        self.children = children

    def to_sql(self):
        return "(" + " AND ".join(child.to_sql() for child in self.children) + ")"

class Or(Expression):
    def __init__(self, *children):
        self.children = children

    def to_sql(self):
        return "(" + " OR ".join(child.to_sql() for child in self.children) + ")"

class QueryBuilder:
    def __init__(self, table, columns=None):
        self.table = table
        self.columns = columns or ["*"]
        self.condition = None

    def where(self, expr):
        self.condition = expr
        return self

    def build(self):
        sql = f"SELECT {', '.join(self.columns)} FROM {self.table}"
        if self.condition:
            sql += f" WHERE {self.condition.to_sql()}"
        return sql + ";"

if __name__ == "__main__":
    cond1 = Condition("age", ">", "30")
    cond2 = Condition("city", "=", "Delhi")
    cond3 = Condition("salary", ">", "50000")

    expression = Or(
        And(cond1, cond2),
        cond3
    )

    qb = QueryBuilder("employees", ["id", "name", "salary"]).where(expression)
    print(qb.build())

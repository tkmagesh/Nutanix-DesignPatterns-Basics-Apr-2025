class QueryBuilder:
    def __init__(self):
        self._select = []
        self._from = ""
        self._where = ""
        self._order_by = ""

    def select(self, *cols):
        self._select = cols
        return self

    def from_table(self, table):
        self._from = table
        return self

    def where(self, condition):
        self._where = condition
        return self

    def order_by(self, column):
        self._order_by = column
        return self

    def build(self):
        query = "SELECT "
        query += ", ".join(self._select) if self._select else "*"
        query += f" FROM {self._from}"
        if self._where:
            query += f" WHERE {self._where}"
        if self._order_by:
            query += f" ORDER BY {self._order_by}"
        return query + ";"

# Client
if __name__ == "__main__":
    query = (
        QueryBuilder()
        .select("id", "name")
        .from_table("users")
        .where("active = 1")
        .order_by("name")
        .build()
    )

    print(query)

class Pizza:
    def __init__(self, dough=None, sauce=None, toppings=None):
        self.dough = dough
        self.sauce = sauce
        self.toppings = toppings or []

    def __repr__(self):
        return f"Pizza(dough='{self.dough}', sauce='{self.sauce}', toppings={self.toppings})"

# Builder
class PizzaBuilder:
    def __init__(self):
        self._pizza = Pizza()

    def set_dough(self, dough):
        self._pizza.dough = dough
        return self

    def set_sauce(self, sauce):
        self._pizza.sauce = sauce
        return self

    def add_topping(self, topping):
        self._pizza.toppings.append(topping)
        return self

    def build(self):
        return self._pizza

# Client
if __name__ == "__main__":
    pizza = (
        PizzaBuilder()
        .set_dough("Thin Crust")
        .set_sauce("Tomato")
        .add_topping("Cheese")
        .add_topping("Olives")
        .build()
    )
    print(pizza)
    """ pizza2 = eval("Pizza(dough='Thin Crust', sauce='Tomato', toppings=['Cheese', 'Olives'])")
    print(pizza2) """
    

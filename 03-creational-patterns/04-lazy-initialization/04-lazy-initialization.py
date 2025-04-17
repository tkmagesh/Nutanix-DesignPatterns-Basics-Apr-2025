class LazyResource:
    def __init__(self):
        self._data = None

    @property
    def data(self):
        if self._data is None:
            print("Initializing data...")
            self._data = self._expensive_operation()
        return self._data

    def _expensive_operation(self):
        return [i * i for i in range(1000000)]

# Usage
resource = LazyResource()
print("Before accessing data")
print(resource.data[:5])  # Data gets initialized here
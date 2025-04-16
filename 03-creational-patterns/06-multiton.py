class RegionalDBConnection:
    _instances = {}

    def __new__(cls, region):
        if region not in cls._instances:
            cls._instances[region] = super().__new__(cls)
            cls._instances[region].region = region
        return cls._instances[region]

    def connect(self):
        return f"Connecting to DB in {self.region}"

# Usage
us_db = RegionalDBConnection("US")
eu_db = RegionalDBConnection("EU")
another_us_db = RegionalDBConnection("US")

print(us_db is another_us_db)  # True
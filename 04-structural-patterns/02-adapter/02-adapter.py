# Target interface
class PaymentProcessor:
    def pay(self, amount):
        raise NotImplementedError()

# Adaptees (Third-party services)
class PayPalService:
    def send_payment(self, amount):
        print(f"PayPal processed payment of {amount}")

class StripeService:
    def charge_card(self, amount):
        print(f"Stripe charged amount {amount}")

class RazorpayService:
    def make_transaction(self, amount):
        print(f"Razorpay processed transaction of {amount}")

# Adapters
class PayPalAdapter(PaymentProcessor):
    def __init__(self, service):
        self.service = service

    def pay(self, amount):
        self.service.send_payment(amount)

class StripeAdapter(PaymentProcessor):
    def __init__(self, service):
        self.service = service

    def pay(self, amount):
        self.service.charge_card(amount)

class RazorpayAdapter(PaymentProcessor):
    def __init__(self, service):
        self.service = service

    def pay(self, amount):
        self.service.make_transaction(amount)

# Client code
def process_payment(processor: PaymentProcessor, amount: int):
    processor.pay(amount)

# Usage
gateways = [
    PayPalAdapter(PayPalService()),
    StripeAdapter(StripeService()),
    RazorpayAdapter(RazorpayService())
]

for gateway in gateways:
    process_payment(gateway, 1000)

# === Common Interface ===
class PaymentProcessor:
    def pay(self, amount):
        raise NotImplementedError()

# === Adaptees ===
class PayPalService:
    def send_payment(self, amount):
        print(f"PayPal: {amount} paid")

class StripeService:
    def charge_card(self, amount):
        print(f"Stripe: {amount} charged")

class RazorpayService:
    def make_transaction(self, amount):
        print(f"Razorpay: {amount} transacted")

# === Adapters ===
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

# === Factory ===
class PaymentGatewayFactory:
    @staticmethod
    def get_processor(gateway: str) -> PaymentProcessor:
        match gateway.lower():
            case "paypal":
                return PayPalAdapter(PayPalService())
            case "stripe":
                return StripeAdapter(StripeService())
            case "razorpay":
                return RazorpayAdapter(RazorpayService())
            case _:
                raise ValueError("Unsupported payment gateway")

# === Client ===
def checkout(amount, gateway_name):
    processor = PaymentGatewayFactory.get_processor(gateway_name)
    processor.pay(amount)

# === Usage ===
checkout(1000, "paypal")
checkout(2000, "stripe")
checkout(3000, "razorpay")

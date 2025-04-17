class Handler:
    def __init__(self, next_handler=None):
        self.next = next_handler

    def handle(self, request):
        if self.next:
            return self.next.handle(request)
        return "No handler could process the request"


class AuthHandler(Handler):
    def handle(self, request):
        if request.get("user") == "admin":
            return "Authenticated"
        return super().handle(request)


class LoggingHandler(Handler):
    def handle(self, request):
        print(f"Logging request: {request}")
        return super().handle(request)


request = {"user": "guest"}
chain = LoggingHandler(AuthHandler())

result = chain.handle(request)
print(result)

class ChatRoomMediator:
    def show_message(self, sender, message):
        pass

class ChatRoom(ChatRoomMediator):
    def __init__(self):
        self.users = []

    def register(self, user):
        self.users.append(user)

    def show_message(self, sender, message):
        for user in self.users:
            if user != sender:
                user.receive(sender.name, message)

class User:
    def __init__(self, name, chatroom):
        self.name = name
        self.chatroom = chatroom
        self.chatroom.register(self)

    def send(self, message):
        print(f"{self.name} sends: {message}")
        self.chatroom.show_message(self, message)

    def receive(self, sender_name, message):
        print(f"{self.name} receives from {sender_name}: {message}")

# Client code
chatroom = ChatRoom()
alice = User("Alice", chatroom)
bob = User("Bob", chatroom)
carol = User("Carol", chatroom)

alice.send("Hi everyone!")
bob.send("Hello Alice!")

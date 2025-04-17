""" Protection Proxy """

from abc import ABC, abstractmethod

class Subject(ABC):
    @abstractmethod
    def read(self):
        pass

class RealDocument(Subject):
    def read(self):
        print("Reading sensitive document")

class DocumentProxy(Subject):
    def __init__(self, user_role):
        self.user_role = user_role
        self.real_doc = RealDocument()

    def read(self):
        if self.user_role != "admin":
            print("Access Denied")
        else:
            self.real_doc.read()

# Usage
doc = DocumentProxy("user")
doc.read()  # Access Denied

admin_doc = DocumentProxy("admin")
admin_doc.read()  # Reading sensitive document
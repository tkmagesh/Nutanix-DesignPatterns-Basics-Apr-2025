from abc import ABC, abstractmethod

class FileSystemComponent(ABC):
    @abstractmethod
    def show(self, indentation=""):
        pass

class File(FileSystemComponent):
    def __init__(self, name):
        self.name = name

    def show(self, indentation=""):
        print(f"{indentation}{self.name}")

class Folder(FileSystemComponent):
    def __init__(self, name):
        self.name = name
        self.children = []

    def add(self, component):
        self.children.append(component)

    def show(self, indentation=""):
        print(f"{indentation}{self.name}/")
        for child in self.children:
            child.show(indentation + "  ")

if __name__ == "__main__":
    root = Folder("root")
    home = Folder("home")
    user = Folder("user")
    file1 = File("file1.txt")
    file2 = File("file2.txt")

    user.add(file1)
    home.add(user)
    root.add(home)
    root.add(file2)

    root.show()

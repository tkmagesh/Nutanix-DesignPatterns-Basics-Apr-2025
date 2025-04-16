from abc import ABC, abstractmethod

# Abstract Products
class Button(ABC):
    @abstractmethod
    def paint(self): pass

class Checkbox(ABC):
    @abstractmethod
    def paint(self): pass

# Concrete Products - Windows
class WindowsButton(Button):
    def paint(self):
        print("Rendering Windows Button")

class WindowsCheckbox(Checkbox):
    def paint(self):
        print("Rendering Windows Checkbox")

# Concrete Products - MacOS
class MacButton(Button):
    def paint(self):
        print("Rendering Mac Button")

class MacCheckbox(Checkbox):
    def paint(self):
        print("Rendering Mac Checkbox")

# Abstract Factory
class GUIFactory(ABC):
    @abstractmethod
    def create_button(self) -> Button: pass

    @abstractmethod
    def create_checkbox(self) -> Checkbox: pass

# Concrete Factories
class WindowsFactory(GUIFactory):
    def create_button(self) -> Button:
        return WindowsButton()
    def create_checkbox(self) -> Checkbox:
        return WindowsCheckbox()

class MacFactory(GUIFactory):
    def create_button(self) -> Button:
        return MacButton()
    def create_checkbox(self) -> Checkbox:
        return MacCheckbox()

# Client
def render_ui(factory: GUIFactory):
    btn = factory.create_button()
    chk = factory.create_checkbox()
    btn.paint()
    chk.paint()

if __name__ == "__main__":
    os = "mac"
    factory = MacFactory() if os == "mac" else WindowsFactory()
    render_ui(factory)

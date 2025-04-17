from typing import List

class Observer:
    def update(self, message: str):
        pass

class Subject:
    def __init__(self):
        self._observers: List[Observer] = []

    def attach(self, observer: Observer):
        self._observers.append(observer)

    def detach(self, observer: Observer):
        self._observers.remove(observer)

    def notify(self, message: str):
        for observer in self._observers:
            observer.update(message)

class NewsAgency(Subject):
    def publish_news(self, news: str):
        print(f"News published: {news}")
        self.notify(news)

class Subscriber(Observer):
    def __init__(self, name: str):
        self.name = name

    def update(self, message: str):
        print(f"{self.name} received news: {message}")

# Usage
agency = NewsAgency()
s1 = Subscriber("Alice")
s2 = Subscriber("Bob")

agency.attach(s1)
agency.attach(s2)

agency.publish_news("Observer Pattern Implemented!")

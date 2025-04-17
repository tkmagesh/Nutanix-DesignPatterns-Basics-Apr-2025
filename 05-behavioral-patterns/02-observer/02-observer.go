package main

import "fmt"

type Observer interface {
	Update(news string)
}

type Subject interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(news string)
}

type NewsAgency struct {
	observers []Observer
}

func (n *NewsAgency) Attach(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *NewsAgency) Detach(o Observer) {
	for i, observer := range n.observers {
		if observer == o {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsAgency) Notify(news string) {
	for _, observer := range n.observers {
		observer.Update(news)
	}
}

func (n *NewsAgency) PublishNews(news string) {
	fmt.Println("News published:", news)
	n.Notify(news)
}

type Subscriber struct {
	name string
}

func (s *Subscriber) Update(news string) {
	fmt.Printf("%s received news: %s\n", s.name, news)
}

func main() {
	agency := &NewsAgency{}
	alice := &Subscriber{name: "Alice"}
	bob := &Subscriber{name: "Bob"}

	agency.Attach(alice)
	agency.Attach(bob)

	agency.PublishNews("Observer Pattern Implemented!")
}

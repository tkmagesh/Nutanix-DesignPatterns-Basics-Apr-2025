package main

import "fmt"

type Handler interface {
	Handle(request map[string]string) string
	SetNext(handler Handler)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) Handle(request map[string]string) string {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return "No handler could process the request"
}

type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(request map[string]string) string {
	if request["user"] == "admin" {
		return "Authenticated"
	}
	return a.BaseHandler.Handle(request)
}

type LoggingHandler struct {
	BaseHandler
}

func (l *LoggingHandler) Handle(request map[string]string) string {
	fmt.Println("Logging request:", request)
	return l.BaseHandler.Handle(request)
}

func main() {
	req := map[string]string{"user": "guest"}

	log := &LoggingHandler{}
	auth := &AuthHandler{}

	log.SetNext(auth)

	fmt.Println(log.Handle(req))
}

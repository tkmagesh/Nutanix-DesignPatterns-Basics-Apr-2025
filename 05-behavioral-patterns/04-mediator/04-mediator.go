package main

import "fmt"

type ChatRoomMediator interface {
	ShowMessage(sender *User, message string)
	Register(user *User)
}

type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) Register(user *User) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) ShowMessage(sender *User, message string) {
	for _, user := range c.users {
		if user != sender {
			user.Receive(sender.name, message)
		}
	}
}

type User struct {
	name     string
	chatroom ChatRoomMediator
}

func NewUser(name string, chatroom ChatRoomMediator) *User {
	user := &User{name: name, chatroom: chatroom}
	chatroom.Register(user)
	return user
}

func (u *User) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.chatroom.ShowMessage(u, message)
}

func (u *User) Receive(senderName, message string) {
	fmt.Printf("%s receives from %s: %s\n", u.name, senderName, message)
}

func main() {
	chatroom := &ChatRoom{}
	alice := NewUser("Alice", chatroom)
	bob := NewUser("Bob", chatroom)
	carol := NewUser("Carol", chatroom)

	alice.Send("Hi everyone!")
	bob.Send("Hey Alice!")
}

package main

import "fmt"

// Command interface
type Command interface {
	Execute()
	Undo()
}

// Receiver
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// Concrete Commands
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

// Invoker
type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) ExecuteCommand(c Command) {
	c.Execute()
	r.history = append(r.history, c)
}

func (r *RemoteControl) UndoLast() {
	if len(r.history) == 0 {
		return
	}
	last := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	last.Undo()
}

// Client
func main() {
	light := &Light{}
	remote := &RemoteControl{}

	onCmd := &LightOnCommand{light}
	offCmd := &LightOffCommand{light}

	remote.ExecuteCommand(onCmd)
	remote.ExecuteCommand(offCmd)
	remote.UndoLast() // Undo LightOff
	remote.UndoLast() // Undo LightOn
}

package main

import "fmt"

type App struct {
	preferredNotificationMedium string
	notifierFactory             INotifierFatory
}

func (app *App) DoWork() {
	msg := "Task Completed"
	var notifier INotifier
	notifier = app.notifierFactory.CreateNotifier(app.preferredNotificationMedium)
	notifier.Send(msg)
}

func NewApp(notificationMedium string, notifierFactory INotifierFatory) *App {
	return &App{
		preferredNotificationMedium: notificationMedium,
		notifierFactory:             notifierFactory,
	}
}

type INotifierFatory interface {
	CreateNotifier(medium string) INotifier
}

type NotifierFactory struct {
}

func (nf *NotifierFactory) CreateNotifier(medium string) INotifier {
	var notifier INotifier
	switch medium {
	case "SMS":
		notifier = NewSMSNotifier("+91-98786-31425")
	case "Email":
		notifier = NewEmailNotifier("tkmagesh77@gmail.com")
	case "Slack":
		notifier = NewSlackNotifier("@tkmagesh")
	}
	return notifier
}

func NewNotifierFactory() *NotifierFactory {
	return &NotifierFactory{}
}

type INotifier interface {
	Send(msg string)
}

type SMSNotifier struct {
	phoneNo string
}

// INotifier interface implementation
func (smsNotifier *SMSNotifier) Send(msg string) {
	fmt.Printf("SMS Notification [%s]: %q", smsNotifier.phoneNo, msg)
}

func NewSMSNotifier(phoneNo string) *SMSNotifier {
	return &SMSNotifier{
		phoneNo: phoneNo,
	}
}

type EmailNotifier struct {
	emailId string
}

// INotifier interface implementation
func (emailNotifier *EmailNotifier) Send(msg string) {
	fmt.Printf("Email Notification [%s]: %q", emailNotifier.emailId, msg)
}

func NewEmailNotifier(emailId string) *EmailNotifier {
	return &EmailNotifier{
		emailId: emailId,
	}
}

type SlackNotifier struct {
	handle string
}

// INotifier interface implementation
func (slackNotifier *SlackNotifier) Send(msg string) {
	fmt.Printf("Slack Notification [%s]: %q", slackNotifier.handle, msg)
}

func NewSlackNotifier(handle string) *SlackNotifier {
	return &SlackNotifier{
		handle: handle,
	}
}

/* Email, SMS, Slack */
func main() {
	notifierFactory := NewNotifierFactory()
	app := NewApp("Slack", notifierFactory)
	app.DoWork()
}

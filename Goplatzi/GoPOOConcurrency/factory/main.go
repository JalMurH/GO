package main

import "fmt"

/*
Programa capaz de manejar diferentes tipos de notificaciones
SMS, Email
*/

type InNotificationFactory interface {
	sendNotification()
	getSender() iSender
}

type iSender interface {
	getSenderMethod() string
	getSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) sendNotification() {
	fmt.Println("Sending Notification SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotification) getSender() iSender {
	return SMSNotificationSender{}
}

func (SMSNotificationSender) getSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) getSenderChannel() string {
	return "Twilo"
}

//Email

type eMailNotification struct {
}

func (eMailNotification) sendNotification() {
	fmt.Println("SEnding Notification Email")
}

type eMailNotificationSender struct {
}

func (eMailNotificationSender) getSenderMethod() string {
	return "Email"
}

func (eMailNotificationSender) getSenderChannel() string {
	return "SES"
}

func (eMailNotification) getSender() iSender {
	return eMailNotificationSender{}
}

func getNotificationFactory(notificacionType string) (InNotificationFactory, error) {
	if notificacionType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificacionType == "Email" {
		return &eMailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f InNotificationFactory) {
	f.sendNotification()
}
func getMethod(f InNotificationFactory) {
	fmt.Println(f.getSender().getSenderChannel())
}
func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	eMailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(eMailFactory)
	getMethod(smsFactory)
	getMethod(eMailFactory)
}

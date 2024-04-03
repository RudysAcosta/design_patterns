package notification

import "fmt"

type INotidicationFactory interface {
	SendNotification()
	GetSender() ISender
}

func GetNotificationFactory(notificationType string) (INotidicationFactory, error) {

	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no notification type")
}

func SendNotification(f INotidicationFactory) {
	f.SendNotification()
}

func GetMethod(f INotidicationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

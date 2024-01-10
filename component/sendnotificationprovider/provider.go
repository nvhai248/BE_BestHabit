package sendnotificationprovider

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// NotificationService represents a service to handle push notifications
type NotificationService struct {
	app    *firebase.App
	ctx    context.Context
	client *messaging.Client
}

// NewNotificationService creates a new instance of NotificationService
func NewNotificationService(ctx context.Context, firebaseConfigPath string) (*NotificationService, error) {
	opt := option.WithCredentialsFile(firebaseConfigPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	return &NotificationService{
		app:    app,
		ctx:    ctx,
		client: client,
	}, nil
}

// SendNotification sends a push notification to the specified device token
func (ns *NotificationService) SendNotification(deviceToken string, title, body string) error {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: deviceToken,
	}

	_, err := ns.client.Send(ns.ctx, message)
	if err != nil {
		return err
	}

	return nil
}

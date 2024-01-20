package main

import (
	"bestHabit/component/sendnotificationprovider"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	projectId := "best-habit"
	privateKeyId := "53db9adca13d306485dbe8c7b1bbac1cada7eb6a"
	privateKey := "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDhvSGT21h1cp4y\nnPlQJ0AZ/eYzhJ6Y1qUM+nX8ItjnbdoyK9B546Mu7Sc9p8I/D5snjpoOMFDYI7C2\nE/2dvd42rvkDZQ6ULjL9J6eGVfdNLb2m1HkV9nmjeD2sYHH3BHwJcrgUCxZkHA5M\nMDtKjgHOxOAoKF8uwO5GHEFv4Lmwn69Fc18SlNOUgweyBAfuKJiq1CrHuQoxVNdh\ni0zO5DZjQ5spzECdWoYKc2FHwO5LrY2zA9zOc5KO3ymoRsJbmu1Xt71bJ/grHZ5E\nIYF1mt+FKHNsW2pTLIwa9Bn87KTBQnfkoCpSB+zcAgWoW3YnMKEz+lXVh2zCa5KV\nayHRg/UbAgMBAAECggEABx3BOzoE9d7hIdNVALPcD3WzlhguL79EgL48jfBDSjhq\nUn6TGeSVNlt/izAIrZg+Webi8GORiN7HKVZD/n8HzE2DyBmddgpmzRg87b00CJjS\ni7nS/2A/wWl++DoqHZFkn5+gMfGPiSRhRwGWPOlxISWPs3RiW8T1PfySC4bYE7tF\n0LFt5set59FsUVL0FemhAip/TFQ/J4si04dqCw2komf4hy98R7JHRq7WyCCZpqM5\nQ6UDpIhRpHrOL+vmg4t/+Lq+unfuImvlZjzYq5YSPc9JBvq8bsiAWRh0+S4Ae89s\nPzqIIsfodAlnriWPocNVKyDWQiE7rxY7soNglr5UQQKBgQDz+QbMJAQA8ThGUjun\nodRsFua/iv2tj6DNUIgaDzS1LeuPkjejKK4anN4jzG416nHMxVJagYCfF5BYzoqv\nS2SWfs3YcqEtVKoiNJ3XhEG2xy+g0HWKG1Wq5jxRnx5I4PPHN0f/QUq2GhmZeSet\nmJ5OtoAh7VEocyo6bte+BXG2UQKBgQDs3f06bv1IqdGNvBAR4Q6rfPD0bZvz1DkS\n59ghV+MXIm85izAQy1eLRygEXOrpDB9EKp+BnQAQXLxFPJrP+OpRXee7H7WOU1l8\nSraaRsFEBGl6WgG0acidGtg9IrQhqU1xpk8T2vTSJgul+4BjSBIVOJULJP3JWayN\n1eaFfGgdqwKBgA/XCEDyyavB/ZRbPHJKyH7oEb036fZ9z8PkyaFfgV2OCLA+nwwh\n1QP3UVjjqfgoK5FO8mTb6Zzqq72IU2rEK1i2DOlTr/FAgPdNkT3v4VBbqFT5k9gO\npEY/QoVOHmo+6LTzeuIwvAgMs8LKIfBca1LS+Ii7XryQlZpLnghBVDuxAoGATdJx\njG1C0kjZDJQpQ3aJ91XJZMVOY8HqLof1vp69gbBSkrlkRWBJlvz97NEKbR8Kdr76\nQP9wMfAF+0l6I7JIagtMQ3Kbl/NShz+U1wNAJDS+4vAHd1r6CoPzX8KzJAwX4ase\neAuMPC87zIDhIb1gE5DPhyXUK/9GbiNE5b6GBpkCgYBs2EAUkdvt5ystYzbaCAlk\ng3QUStuD8Wet2GeZZ6D2lyWL+4W1hFjGCz3tyM/qR/gC9qzNqqfRNTpwFyvKK7Tl\n7DxFKkD7JEcbkOgvzq1IZcLzBfbggLkVcmjtANdNWfIW2c9hCx4uwsL30mHJM3YL\nJ4L0Y+Kk3GXaCQ4E0iK5fQ==\n-----END PRIVATE KEY-----\n"
	clientEmail := "firebase-adminsdk-c3by4@best-habit.iam.gserviceaccount.com"
	clientId := "101675413047912112311"

	ctx := context.Background()

	notificationService, err := sendnotificationprovider.NewNotificationService(ctx, projectId, privateKeyId, privateKey, clientEmail, clientId)
	if err != nil {
		log.Fatalf("Error creating NotificationService: %v", err)
	}

	deviceToken := "your-device-token"
	title := "Test Notification"
	body := "This is a test notification."

	err = notificationService.SendNotification(deviceToken, title, body)
	if err != nil {
		log.Fatalf("Error sending notification: %v", err)
	}

	fmt.Println("Notification sent successfully!")

	time.Sleep(5 * time.Second)
}

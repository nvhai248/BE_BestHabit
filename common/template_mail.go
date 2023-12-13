package common

import (
	"fmt"
	"os"
)

type Email struct {
	Subject     string
	Content     string
	To          []string
	Cc          []string
	Bcc         []string
	AttachFiles []string
}

func NewEmailVerifyAccount(to []string, token string) *Email {
	return &Email{
		Subject: "Verify your account",
		Content: fmt.Sprintf(`
			<div style="font-family: Arial, sans-serif; max-width: 600px; margin: auto;">
				<h1 style="color: #333; text-align: center;">Welcome!</h1>
				<p style="font-size: 16px; color: #555;">Hi there,</p>
				<p style="font-size: 16px; color: #555;">Thank you for signing up! To complete the verification process, please click the button below.</p>
				<div style="text-align: center; margin-top: 20px;">
				<a href="%s/verification?token_id=%s" style="display: inline-block; padding: 12px 24px; font-size: 16px; text-decoration: none; background-color: #007bff; color: #fff; border-radius: 5px;">Verify Email</a>
				</div>
				<p style="font-size: 14px; color: #999; text-align: center; margin-top: 20px;">Note: This link will expire in 1 day.</p>
			</div>
		`, os.Getenv("DOMAIN_CLIENT"), token),
		To:          to,
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
}

func NewRequireResetPw(to []string, token string) *Email {
	return &Email{
		Subject: "Password Reset Request",
		Content: fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>Password Reset Request</title>
		</head>
		<body style="font-family: Arial, sans-serif; max-width: 600px; margin: auto; color: #333;">
			<div style="text-align: center;">
				<h1 style="color: #333;">Password Reset Request</h1>
				<p style="font-size: 16px; color: #555;">Hi there,</p>
				<p style="font-size: 16px; color: #555;">You have requested to reset your password. Please click the button below to proceed:</p>
				<div style="text-align: center; margin-top: 20px;">
					<a href="%s/resetPw?token_id=%s" style="display: inline-block; padding: 12px 24px; font-size: 16px; text-decoration: none; background-color: #007bff; color: #fff; border-radius: 5px;">Reset Password</a>
				</div>
				<p style="font-size: 14px; color: #999; margin-top: 20px;">Please note: This link will expire in 1 day.</p>
			</div>
		</body>
		</html>
		`, os.Getenv("DOMAIN_CLIENT"), token),
		To:          to,
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
}

func NewRequireResetPwAfterChangePass(to []string, token string) *Email {
	return &Email{
		Subject: "Password Reset Request",
		Content: fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: auto;">
			<h1 style="color: #333; text-align: center;">Password Reset Request</h1>
			<p style="font-size: 16px; color: #555; text-align: center;">Hi there,</p>
			<p style="font-size: 16px; color: #555;">We have noticed that your password has recently been changed. If this wasn't done by you, please reset your password by clicking the button below:</p>
			<div style="text-align: center; margin-top: 20px;">
			<a href="%s/resetPw?token_id=%s" style="display: inline-block; padding: 12px 24px; font-size: 16px; text-decoration: none; background-color: #007bff; color: #fff; border-radius: 5px;">Reset Password</a>
			</div>
			<p style="font-size: 14px; color: #999; text-align: center; margin-top: 20px;">If you did not request this change, please reset your password immediately.</p>
			<p style="font-size: 14px; color: #999; text-align: center; margin-top: 20px;">Please note: This link will expire in 1 day.</p>
		</div>
		`, os.Getenv("DOMAIN_CLIENT"), token),
		To:          to,
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
}

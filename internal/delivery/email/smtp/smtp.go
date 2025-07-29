package smtp

import smtpPkg "net/smtp"

type EmailDeliveryClient struct {
	smtpSever  string
	agentEmail string
	auth       smtpPkg.Auth
}

func NewEmailDeliveryClient(smtpServer, agentEmail string, smtpAuth smtpPkg.Auth) *EmailDeliveryClient {
	return &EmailDeliveryClient{
		smtpSever:  smtpServer,
		agentEmail: agentEmail,
		auth:       smtpAuth,
	}
}

func (c *EmailDeliveryClient) SendEmail(to string, subject string, body string) error {
	// Implementation for sending email via SMTP server
	// This is a placeholder for actual SMTP logic
	return smtpPkg.SendMail(c.smtpSever, c.auth, c.agentEmail, []string{to}, []byte(
		"Subject: "+subject+"\r\n"+
			"Content-Type: text/plain; charset=UTF-8\r\n"+
			"\r\n"+
			body,
	))
}

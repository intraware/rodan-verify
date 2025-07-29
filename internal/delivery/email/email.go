package email

type EmailDelivery interface {
	SendEmail(to string, subject string, body string) error
}

type EmailDeliveryAgent int

const (
	_                                EmailDeliveryAgent = iota
	EmailDeliveryAgentSMTP                              // 1
	EmailDeliveryAgentMicrosoftGraph                    // 2
)

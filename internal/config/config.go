package config

import (
	"os"

	"github.com/intraware/rodan-verify/internal/delivery/email"
	"github.com/intraware/rodan-verify/internal/otp"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server         ServerConfig         `yaml:"server"`
	Otp            OtpConfig            `yaml:"otp"`
	Email          EmailConfig          `yaml:"email"`
	Smtp           SMTPConfig           `yaml:"smtp"`
	MicrosoftGraph MicrosoftGraphConfig `yaml:"microsoft_graph"`
	Validation     ValidationConfig     `yaml:"validation"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type OtpConfig struct {
	Format otp.OTP_FORMAT `yaml:"format"` // Options: numeric (1), alphanumeric (2), alphabetic (3)
	Length int            `yaml:"length"`
	Expiry int64          `yaml:"expiry"` // in seconds
}

type EmailConfig struct {
	AgentEmail string                   `yaml:"agent_email"`
	Provider   email.EmailDeliveryAgent `yaml:"provider"` // Options: SMTP (1), Microsoft Graph (2)
}

type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MicrosoftGraphConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	TenantID     string `yaml:"tenant_id"`
}

type ValidationConfig struct {
	AllowedEmailRegex string `yaml:"allowed_email_regex"` // Regex pattern to validate email addresses
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

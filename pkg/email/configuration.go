package email

import (
	"github.com/arehmandev/Log-mailer/pkg/utils"
)

//Config - a basic structure of configurations
type Config struct {
	From        Fields            `json:"from"`
	To          Fields            `json:"to"`
	Subject     string            `json:"subject"`
	Server      string            `json:"server"`
	Port        string            `json:"port"`
	Credentials Credentials       `json:"credentials"`
	Logs        string            `json:"logs"`
	Interval    string            `json:"interval"`
	Reset       string            `json:"reset"`
	Message     string            `json:"message"`
	Headers     map[string]string `json:"headers"`
}

//Fields - creates a basic email structure
type Fields struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Credentials - used for logging into the email account of sender
type Credentials struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

// GenerateConfig -
func GenerateConfig() *Config {

	newConfig := &Config{
		From: Fields{
			Name:  utils.Ask("Enter \"From\" name:\t"),
			Email: utils.Ask("Enter \"From\" email:\t"),
		},
		To: Fields{
			Name:  utils.Ask("Enter \"To\" name:\t"),
			Email: utils.Ask("Enter \"To\" email:\t"),
		},
		Subject: utils.Ask("Enter subject:\t\t"),
		Server:  utils.Ask("Enter SMTP server:\t"),
		Port:    utils.Ask("Port:\t\t\t"),
		Credentials: Credentials{
			Username: utils.Ask("Username:\t\t"),
			Password: utils.Ask("Password:\t\t"),
		},
		Logs:     utils.Ask("Location of logs:\t"),
		Interval: utils.Ask("Interval:\t\t"),
		Reset:    utils.Ask("Reset log file? (y/n):\t"),
	}

	if newConfig.Reset != "y" {
		newConfig.Reset = "false"
	} else {
		newConfig.Reset = "true"
	}

	return newConfig

}

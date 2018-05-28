package main

import (
	"bufio"
	"fmt"
	"os"
)

//Config - a basic structure of configurations
type Config struct {
	From        EmailStruct `json:"from"`
	To          EmailStruct `json:"to"`
	Subject     string      `json:"subject"`
	Server      string      `json:"server"`
	Port        string      `json:"port"`
	Credentials Credentials `json:"credentials"`
	Logs        string      `json:"logs"`
	Interval    string      `json:"interval"`
	Reset       string      `json:"reset"`
	Message     string
	Headers     map[string]string
}

//EmailStruct - creates a basic email structure
type EmailStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Credentials - used for logging into the email account of sender
type Credentials struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func (c *Config) generateConfig() {

	c = &Config{
		From: EmailStruct{
			Name:  ask("Enter \"From\" name:\t"),
			Email: ask("Enter \"From\" email:\t"),
		},
		To: EmailStruct{
			Name:  ask("Enter \"To\" name:\t"),
			Email: ask("Enter \"To\" email:\t"),
		},
		Subject: ask("Enter subject:\t\t"),
		Server:  ask("Enter SMTP server:\t"),
		Port:    ask("Port:\t\t\t"),
		Credentials: Credentials{
			Username: ask("Username:\t\t"),
			Password: ask("Password:\t\t"),
		},
		Logs:     ask("Location of logs:\t"),
		Interval: ask("Interval:\t\t"),
		Reset:    ask("Reset log file? (y/n):\t"),
	}
	if c.Reset != "y" {
		c.Reset = "false"
	} else {
		c.Reset = "true"
	}

}

func ask(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	r, _, _ := reader.ReadLine()
	return string(r)
}

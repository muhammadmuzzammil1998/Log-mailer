package main

//Config is a basic structure of configurations
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
}

//EmailStruct creates a basic email address including Name and Email ID
type EmailStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Credentials are used for loging into the email account of sender
type Credentials struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

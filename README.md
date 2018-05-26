# Log mailer [![CircleCI](https://circleci.com/gh/muhammadmuzzammil1998/Log-mailer.svg?style=svg)](https://circleci.com/gh/muhammadmuzzammil1998/Log-mailer)
Log mailer is a program I made to email log files from a server to me so I don't have to manually check logs everytime. It uses GoLang's `"net/smtp"` to email.

## Build
Prerequisites: [Git](https://git-scm.com/downloads), [Go](https://golang.org/dl/), SMTP server credentials.
```bash
$ git clone https://github.com/muhammadmuzzammil1998/Log-mailer
$ cd Log-mailer
$ go build
```

[![asciicast](https://asciinema.org/a/DtiFec3m4gAnIKzjo2Qw9WH9l.png)](https://asciinema.org/a/DtiFec3m4gAnIKzjo2Qw9WH9l)

## Generating configuration file.
```bash
$ ./Log-mailer -generate
Location to store configuration file (default: ./configuration.json):
Enter "From" name:      Muhammad Muzzammil
Enter "From" email:     email@muzzammil.xyz
Enter "To" name:        Muhammad Muzzammil
Enter "To" email:       muhammadmuzzammil.cs@gmail.com
Enter subject:          Error logs from server
Enter SMTP server:      smtp.server.lnk
Port:                   25
Username:               email@muzzammil.xyz
Password:               r34lly53cur3p455w0rd
Location of logs:       /path/to/logs.log
Interval:               6h
Reset log file? (y/n):  y
Configuration file generated: configuration.json.
```

[![asciicast](https://asciinema.org/a/rBWW7nvmJtezsRnn3dzfvafvN.png?AAAAAAH)](https://asciinema.org/a/rBWW7nvmJtezsRnn3dzfvafvN)

### Description
| Input              | Description                                                                                 |
|:-------------------|:--------------------------------------------------------------------------------------------|
| Configuration File | A single file to store data. Default is `./configuration.json` in current directory.        |
| "From" data        | The author of the message. Used for "From" header                                           |
| "To" data          | The address of the primary recipient of the message. Used for "To" Header                   |
| "Subject"          | What you want the subject to be. Used for "Subject" header.                                 |
| SMTP Server        | Link to your mail server.                                                                   |
| Port               | Port to use for the server.                                                                 |
| Username           | Username for SMTP server                                                                    |
| Password           | Password for SMTP server                                                                    |
| Location           | Location of the log file which is to be emailed                                             |
| Interval           | The duration in which email is repeated. Valid units are `"ns", "us", "ms", "s", "m", "h"`. |
| Reset logfile      | Empties the log file after a successful email.                                              |

### JSON structure of configuration file
```json
{
    "from": {
        "name": "Muhammad Muzzammil",
        "email": "email@muzzammil.xyz"
    },
    "to": {
        "name": "Muhammad Muzzammil",
        "email": "muhammadmuzzammil.cs@gmail.com"
    },
    "subject": "Error logs from server",
    "server": "smtp.server.lnk",
    "port": "25",
    "credentials": {
        "user": "email@muzzammil.xyz",
        "password": "r34lly53cur3p455w0rd"
    },
    "logs": "/path/to/logs.log",
    "interval": "6h",
    "reset": "true"
}
```

### Go structure of configurations
```go
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

type EmailStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Credentials struct {
	Username string `json:"user"`
	Password string `json:"password"`
}
```

## Generating an empty configuration file.
Just in case someone doesn't want to do it interactively.
```bash
$ .\Log-mailer.exe -generate -empty
Location to store configuration file (default: ./configuration.json): configuration_template.json
Empty configuration file generated: configuration_template.json.
```

## Starting Log mailer
### Using default (configuration.json) configuration
```bash
$ ./Log-mailer
```
### Using different configurations
```bash
$ ./Log-mailer -conf [filename]
```
You can start it as different background processes with `-conf` flag and different configurations for multiple log files.

## Help
```
Usage of ./Log-mailer:
  -conf string
        Configuration file to load. (default "configuration.json")
  -empty
        Generate an empty configuration file. Use with -generate.
  -generate
        Generate configuration file interactively.
```

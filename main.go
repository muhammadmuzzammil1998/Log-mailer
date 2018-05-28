package main

import (
	"flag"
	"log"
	"os"

	"github.com/arehmandev/Log-mailer/pkg/email"
)

func main() {
	var createJSON, emptyJSON bool
	var fileNameJSON string
	flag.BoolVar(&createJSON, "generate", false, "Generate configuration file interactively.")
	flag.BoolVar(&emptyJSON, "empty", false, "Generate an empty configuration file.")
	flag.StringVar(&fileNameJSON, "conf", "configuration.json", "Configuration file to load.")
	flag.Parse()

	if _, err := os.Stat(fileNameJSON); err != nil {
		log.Fatalf("Unable to find configuration file (%s).\n", fileNameJSON)
	}

	email := new(email.Config)

	email.GenerateConfigJSONFiles(createJSON, emptyJSON, fileNameJSON)
	email.Email(fileNameJSON)

}

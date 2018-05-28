package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/arehmandev/Log-mailer/pkg/config"
	"github.com/arehmandev/Log-mailer/pkg/email"
	"github.com/arehmandev/Log-mailer/pkg/utils"
)

func main() {
	var createJSON, emptyJSON bool
	var fileNameJSON string
	flag.BoolVar(&createJSON, "generate", false, "Generate configuration file interactively.")
	flag.BoolVar(&emptyJSON, "empty", false, "Generate an empty configuration file. Use with -generate.")
	flag.StringVar(&fileNameJSON, "conf", "configuration.json", "Configuration file to load.")
	flag.Parse()

	// Generate an empty JSON file is empty flag was provided
	config.GenerateEmptyJSON(emptyJSON, fileNameJSON)

	// Generate the JSON file if generate flag was given
	config.GenerateJSON(createJSON, fileNameJSON)

	if _, err := os.Stat(fileNameJSON); err != nil {
		log.Fatalf("Unable to find configuration file (%s).\n", fileNameJSON)
	}

	data, err := ioutil.ReadFile(fileNameJSON)
	utils.Check(err)

	c := new(email.Config)
	err = json.Unmarshal(data, c)
	utils.Check(err)
	utils.Repeat(c.EmailLogs, c.Interval)
}

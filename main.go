package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var createJSON, emptyJSON bool
	var fileNameJSON string
	flag.BoolVar(&createJSON, "generate", false, "Generate configuration file interactively.")
	flag.BoolVar(&emptyJSON, "empty", false, "Generate an empty configuration file. Use with -generate.")
	flag.StringVar(&fileNameJSON, "conf", "configuration.json", "Configuration file to load.")
	flag.Parse()

	// Generate an empty JSON file is empty flag was provided
	generateEmptyJSON(emptyJSON, fileNameJSON)

	// Generate the JSON file if generate flag was given
	generateJSON(createJSON, fileNameJSON)

	if _, err := os.Stat(fileNameJSON); err != nil {
		log.Fatalf("Unable to find configuration file (%s).\n", fileNameJSON)
	}

	data, err := ioutil.ReadFile(fileNameJSON)
	check(err)

	c := new(Config)
	err = json.Unmarshal(data, c)
	check(err)

	c.generateConfig()

	repeat(c.emailLogs, c.Interval)
}

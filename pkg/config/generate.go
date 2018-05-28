package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arehmandev/Log-mailer/pkg/email"
	"github.com/arehmandev/Log-mailer/pkg/utils"
)

// GenerateJSON -
func GenerateJSON(createJSON bool, fileNameJSON string) {

	if !createJSON {
		return
	}

	checkForOverwrite(fileNameJSON)

	f, err := os.OpenFile(fileNameJSON, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	c := new(email.Config)
	c.GenerateConfig()

	j, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "%s", j)
	fmt.Printf("Configuration file generated: %s.\n", fileNameJSON)
	return

}

// GenerateEmptyJSON -
func GenerateEmptyJSON(emptyJSON bool, fileNameJSON string) {

	if !emptyJSON {
		return
	}

	f, err := os.OpenFile(fileNameJSON, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	j, err := json.MarshalIndent(&email.Config{}, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "%s", j)
	fmt.Printf("Empty configuration file generated: %s.\n", fileNameJSON)
	return

}

func checkForOverwrite(fileNameJSON string) {

	reader := bufio.NewReader(os.Stdin)
	if r := utils.Ask("Location to store configuration file (default: ./" + fileNameJSON + "): "); strings.TrimSpace(r) != "" {
		fileNameJSON = r
	}
	if fStat, err := os.Stat(fileNameJSON); err == nil && !fStat.IsDir() {
		fmt.Printf("Configuration file (%s) exists, overwrite? (y/n): ", fileNameJSON)
		if r, _ := reader.ReadString('\n'); strings.ToLower(strings.TrimSpace(r)) != "y" {
			return
		}
		err := os.Remove(fileNameJSON)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

package email

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arehmandev/Log-mailer/pkg/utils"
)

// GenerateJSON -
func GenerateJSON(createJSON, emptyJSON bool, fileNameJSON string) {

	if !createJSON {
		fmt.Println("[SKIPPING] No configuration.json was created")
		return
	}

	checkForOverwrite(fileNameJSON)

	f, err := os.OpenFile(fileNameJSON, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	c := GenerateConfig()

	fmt.Println(c)

	j, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "%s", j)
	fmt.Printf("Configuration file generated: %s.\n", fileNameJSON)
	os.Exit(0)

}

// GenerateEmptyJSON -
func GenerateEmptyJSON(emptyJSON bool, fileNameJSON string) {

	if !emptyJSON {
		return
	}

	err := os.RemoveAll(fileNameJSON)
	utils.Check(err)

	f, err := os.OpenFile(fileNameJSON, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	j, err := json.MarshalIndent(&Config{}, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "%s", j)
	fmt.Printf("Empty configuration file generated: %s.\n", fileNameJSON)
	os.Exit(0)

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

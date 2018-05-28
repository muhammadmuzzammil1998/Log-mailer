package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func generateJSON(createJSON bool, fileNameJSON string) {

	if !createJSON {
		return
	}

	checkForOverwrite(fileNameJSON)

	f, err := os.OpenFile(fileNameJSON, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	c := new(Config)
	c.generateConfig()

	j, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(f, "%s", j)
	fmt.Printf("Configuration file generated: %s.\n", fileNameJSON)
	return

}

func generateEmptyJSON(emptyJSON bool, fileNameJSON string) {

	if !emptyJSON {
		return
	}

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
	return

}

func checkForOverwrite(fileNameJSON string) {

	reader := bufio.NewReader(os.Stdin)
	if r := ask("Location to store configuration file (default: ./" + fileNameJSON + "): "); strings.TrimSpace(r) != "" {
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

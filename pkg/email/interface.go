package email

import (
	"encoding/json"
	"io/ioutil"

	"github.com/arehmandev/Log-mailer/pkg/utils"
)

// Emailer -
type Emailer interface {
	GenerateConfigJSONFiles(bool, bool, string)
	Email(string)
}

// GenerateConfigJSONFiles -
func (c *Config) GenerateConfigJSONFiles(createJSON, emptyJSON bool, fileNameJSON string) {

	// Generate an empty JSON file is empty flag was provided
	GenerateEmptyJSON(emptyJSON, fileNameJSON)

	// Generate the JSON file if generate flag was given
	GenerateJSON(createJSON, emptyJSON, fileNameJSON)

}

// Email -
func (c *Config) Email(fileNameJSON string) {

	data, err := ioutil.ReadFile(fileNameJSON)
	utils.Check(err)

	err = json.Unmarshal(data, c)
	utils.Check(err)
	utils.Repeat(c.EmailLogs, c.Interval)
}

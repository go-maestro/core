package maestro

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// ParseDefinition instantiate a Definition from definitons file
func ParseDefinition(filePath string) (*Definition, error) {
	definition := Definition{Version: "1.0"}
	fileContent, err := ioutil.ReadFile(filePath)

	// its used to set the paths for steps
	definition.filePath = filePath
	definition.fileDir = filepath.Dir(filePath)

	if err != nil {
		return nil,
			fmt.Errorf("Error reading definitions: %s", err)
	}

	err = yaml.Unmarshal(fileContent, &definition)

	if err != nil {
		return nil,
			fmt.Errorf("Error parsing definitions: %s", err)
	}

	if err := definition.validate(); err != nil {
		return nil, fmt.Errorf("Errors validating file: %s", err)
	}

	return &definition, nil
}

package maestro

import (
	"fmt"
	"os"
)

// Step represents a step
type Step struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`

	DependsOn []string `yaml:"depends_on"`
	Files     []string `yaml:"files"`
}

// Validate verify if step has a valid declaration
func (s Step) Validate(d *Definition) error {
	if err := s.validateType(); err != nil {
		return err
	}

	if err := s.validateFilesExists(d); err != nil {
		return err
	}

	return nil
}

func (s Step) validateType() error {
	errMsg := fmt.Sprintf("invalid type %s", s.Type)

	switch s.Type {
	case "KUBERNETES":
		return nil

	default:
		return fmt.Errorf(
			"steps.%s.type: %s", s.Name, errMsg,
		)
	}
}

func (s Step) validateFilesExists(d *Definition) error {
	for fileIndex, fileName := range s.Files {
		filePath := fmt.Sprintf("%s/%s", d.fileDir, fileName)

		if _, err := os.Stat(filePath); err != nil {
			return fmt.Errorf(
				"steps.%s.files.%d: %s", s.Name, fileIndex, err,
			)
		}
	}

	return nil
}

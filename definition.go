package maestro

import (
	"fmt"
	"strings"
)

// Definition represents a definition of pipeline
type Definition struct {
	Name    string `yaml:"name"`
	Steps   []Step `yaml:"steps"`
	Version string `yaml:"version"`

	fileDir  string
	filePath string
}

// validate verify the syntax of a definitions file
func (d *Definition) validate() (err error) {
	// Use the same step.Name in more one step
	err = d.validateStepNameUniqueness()
	if err != nil {
		return err
	}

	// Some of step.File points to a inexistent file
	err = d.validateEachStep()
	if err != nil {
		return err
	}

	// For now, ensure only the version 1.0 of definition
	if d.Version != "1.0" {
		return fmt.Errorf("Invalid version %s", d.Version)
	}

	return nil
}

// validateStepNameUniqueness verify if Step.name is unique
func (d *Definition) validateStepNameUniqueness() error {
	var stepNames []string
	var repeteadNames []string

	if d.Version != "1.0" {
		return fmt.Errorf("Invalid version %s", d.Version)
	}

	for _, step := range d.Steps {
		for _, stepName := range stepNames {
			if step.Name == stepName {
				repeteadNames = append(repeteadNames, stepName)
			}
		}

		stepNames = append(stepNames, step.Name)
	}

	if len(repeteadNames) != 0 {
		return fmt.Errorf(
			"Step names must be unique, but %s appears multiple times",
			strings.Join(repeteadNames, ","),
		)
	}

	return nil
}

// validateStep call validations for-each step
func (d *Definition) validateEachStep() error {
	for _, step := range d.Steps {
		if err := step.Validate(d); err != nil {
			return err
		}
	}

	return nil
}

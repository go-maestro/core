package maestro_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	maestro "github.com/go-maestro/core"
	"github.com/stretchr/testify/assert"
)

func TestParseDefinition(t *testing.T) {
	pkgPath, _ := os.Getwd()
	basePath := fmt.Sprintf("%s/data/definitions/", pkgPath)

	t.Run("when definition file doesn't exists", func(t *testing.T) {
		filePath := "/inexistent/file.yaml"
		_, err := maestro.ParseDefinition(filePath)

		t.Run("should return an error", func(t *testing.T) {
			assert.Errorf(
				t, err, "Error reading definitions: open %s: no such file or directory", filePath,
			)
		})
	})

	t.Run("when definition is a invalid yaml", func(t *testing.T) {
		filePath := fmt.Sprintf("%s/invalid.yaml", basePath)
		_, err := maestro.ParseDefinition(filePath)

		t.Run("should return an error", func(t *testing.T) {
			assert.Errorf(
				t, err, "Error parsing definitions: content is not a yaml",
			)
		})
	})

	t.Run("when definitions have repeat step names", func(t *testing.T) {
		filePath := fmt.Sprintf("%s/repeated.steps.yaml", basePath)
		_, err := maestro.ParseDefinition(filePath)

		t.Run("should return an error", func(t *testing.T) {
			msg := fmt.Sprintf("Errors validating file: Step names must be unique, but %s appears multiple times", "step1")

			assert.EqualError(t, err, msg)
		})
	})

	t.Run("when definition doesn't declare a version", func(t *testing.T) {
		filePath := fmt.Sprintf("%s/without.version.yaml", basePath)
		definition, err := maestro.ParseDefinition(filePath)

		t.Run("should defaulty version to 1.0", func(t *testing.T) {
			assert.Equal(t, definition.Version, "1.0")
		})

		t.Run("shouldn't return error", func(t *testing.T) {
			assert.Nil(t, err)
		})
	})

	t.Run("when Step timeout was not declared", func(t *testing.T) {
		t.Run("should defaulty to 2 minutes", func(t *testing.T) {})
	})

	t.Run("when definition is valid", func(t *testing.T) {
		t.Run("when one file of Step files doesn't exists", func(t *testing.T) {
			filePath := fmt.Sprintf("%s/valid/inexistent.file.yaml", basePath)
			_, err := maestro.ParseDefinition(filePath)

			t.Run("should return an error", func(t *testing.T) {
				regex := "Errors validating file: steps\\.\\w+\\.files.\\d+: stat (.*): no such file or directory(.*)"

				assert.Regexp(t, regexp.MustCompile(regex), err)
			})
		})

		t.Run("when all Step files exists", func(t *testing.T) {
			filePath := fmt.Sprintf("%s/valid/everything.ok.yaml", basePath)
			def, err := maestro.ParseDefinition(filePath)

			t.Run("should return a Definition", func(t *testing.T) {
				assert.IsType(t, &maestro.Definition{}, def)
			})

			t.Run("shouldn't return an error", func(t *testing.T) {
				assert.Nil(t, err)
			})

			t.Run("should set Name property", func(t *testing.T) {
				assert.NotNil(t, def.Name)
			})

			t.Run("should set Steps property", func(t *testing.T) {
				assert.IsType(t, []maestro.Step{}, def.Steps)
			})
		})
	})
}

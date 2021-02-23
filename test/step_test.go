package maestro_test

import (
	"regexp"
	"testing"

	maestro "github.com/go-maestro/core"
	"github.com/stretchr/testify/assert"
)

func TestStepValidate(t *testing.T) {
	definition := maestro.Definition{
		Name:    "MyDefinition",
		Version: "1.0",
		Steps:   []maestro.Step{},
	}

	t.Run("when everything is ok", func(t *testing.T) {
		step := maestro.Step{
			Name:  "MyStep",
			Type:  "KUBERNETES",
			Files: []string{""},
		}

		t.Run("shouldn't return an error", func(t *testing.T) {
			assert.Nil(t, step.Validate(&definition))
		})
	})

	t.Run("when Type is not valid", func(t *testing.T) {
		step := maestro.Step{
			Name:  "MyStep",
			Type:  "BLABLABLA",
			Files: []string{""},
		}

		t.Run("should return an error", func(t *testing.T) {
			assert.Errorf(t,
				step.Validate(&definition),
				"steps.MyStep.type: invalid type %s",
				step.Type,
			)
		})
	})

	t.Run("when Name have spaces", func(t *testing.T) {
		t.Run("should return an error", func(t *testing.T) {
			step := maestro.Step{
				Name:  "My Step",
				Type:  "BLABLABLA",
				Files: []string{""},
			}

			t.Run("should return an error", func(t *testing.T) {
				assert.Errorf(t,
					step.Validate(&definition),
					"steps.MyStep.name: expected string without spaces, got %s",
					step.Name,
				)
			})
		})
	})

	t.Run("when Files doesn't exists", func(t *testing.T) {
		step := maestro.Step{
			Name: "MyStep",
			Type: "KUBERNETES",
			Files: []string{
				"it/is/a/inexistent/path.yaml",
			},
		}

		t.Run("should return an error", func(t *testing.T) {
			err := step.Validate(&definition)
			regex := "steps.\\w+.files.\\d+: stat (.*): no such file or directory$"

			assert.Regexp(t, regexp.MustCompile(regex), err)
		})
	})
}

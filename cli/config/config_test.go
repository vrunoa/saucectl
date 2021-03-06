package config

import (
	"testing"

	"github.com/docker/docker/pkg/testutil/assert"
	"gotest.tools/v3/fs"
)

func TestJobConfiguration(t *testing.T) {
	dir := fs.NewDir(t, "fixtures",
		fs.WithFile("invalid_config.yaml", "foo", fs.WithMode(0755)),
		fs.WithFile("valid_config.yaml", "apiVersion: 1.2", fs.WithMode(0755)))
	defer dir.Remove()

	cases := []struct {
		Name       string
		Input      string
		ShouldPass bool
	}{
		{"With nil filename", "", false},
		{"With non existing config", "/dont/exist", false},
		{"With non invalid config", dir.Path() + "/invalid_config.yaml", false},
		{"With valid config", dir.Path() + "/valid_config.yaml", true},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			configObject, err := NewJobConfiguration(tc.Input)
			if !tc.ShouldPass {
				if err == nil {
					t.Error("No error was returned for failing test case")
				}

				return
			}

			if err != nil {
				t.Error("Error was returned for passing test case")
			}

			assert.Equal(t, configObject.APIVersion, "1.2")
		})
	}
}

func TestRunnerConfiguration(t *testing.T) {
	dir := fs.NewDir(t, "fixtures",
		fs.WithFile("valid_config.yaml", "rootDir: /foo/bar", fs.WithMode(0755)))
	defer dir.Remove()

	configObject, err := NewRunnerConfiguration(dir.Path() + "/valid_config.yaml")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, configObject.RootDir, "/foo/bar")
}

package manager

import (
	"gopkg.in/yaml.v2"
	"os"
	"strings"
	"testing"
)

func TestConfig_GenerateFile(t *testing.T) {
	t.Run("Generated file is not empty", func(t *testing.T) {
		var cfg = Config{}
		if content := cfg.GenerateFile(); string(content) == "" {
			t.Errorf("File is empty.")
		}
	})

	t.Run("Generated file can be turned into an config object", func(t *testing.T) {
		var cfg = Config{}
		var content = cfg.GenerateFile()
		var generatedCfg = Config{}

		if err := yaml.Unmarshal(content, &generatedCfg); err != nil {
			t.Errorf("Generated yaml could not be unmarshalled. (%s)", err.Error())
		}

		var dependencyLength = len(generatedCfg.Dependencies)
		if dependencyLength > 0 {
			t.Errorf("The generated config should contain 0 dependencies. %v found.", dependencyLength)
		}
	})

	t.Run("Generated File contains existing data", func(t *testing.T) {
		var cfg = Config{}
		var dep1 = Dependency{
			Target: "foo",
			Source: "bar",
			Type:   "someType",
		}
		var dep2 = Dependency{
			Target: "john",
			Source: "doe",
			Type:   "someOtherType",
		}

		cfg.Dependencies = append(cfg.Dependencies, &dep1)
		cfg.Dependencies = append(cfg.Dependencies, &dep2)

		var content = cfg.GenerateFile()
		var generatedCfg = Config{}

		if err := yaml.Unmarshal(content, &generatedCfg); err != nil {
			t.Errorf("Generated yaml could not be unmarshalled. (%s)", err.Error())
		}

		for idx, dep := range cfg.Dependencies {
			var genDep = generatedCfg.Dependencies[idx]

			if dep.Target != genDep.Target {
				t.Errorf("Target does not match. Expected: %s, Got %s", dep.Target, genDep.Target)
			}
			if dep.Source != genDep.Source {
				t.Errorf("Source does not match. Expected: %s, Got %s", dep.Source, genDep.Source)
			}
			if dep.Type != genDep.Type {
				t.Errorf("Type does not match. Expected: %s, Got %s", dep.Type, genDep.Type)
			}
		}
	})
}

func TestConfig_WriteFile(t *testing.T) {
	t.Run("File is being created", func(t *testing.T) {
		var c = Config{}
		var fileName = "../test/.edm.test.yaml"

		if err := c.WriteFile(fileName); err != nil {
			t.Errorf("File %s could not be written: %s", fileName, err)
		}

		fileInfo, err := os.Stat(fileName)
		if err != nil {
			t.Errorf("%s", err)
			return
		}

		if fileInfo.Size() == 0 {
			t.Errorf("File %s is empty, but shouldn't be.", fileName)
		}

		if err := os.Remove(fileName); err != nil {
			t.Errorf("File %s could not be removed: %s", fileName, err)
		}
	})
}

func TestOpenConfig(t *testing.T) {
	t.Run("Open saved configuration.", func(t *testing.T) {
		t.Run("Fails on opening non-existent file", func(t *testing.T) {
			_, err := OpenConfig("something/non/existing.jar")

			if err == nil {
				t.Errorf("This should have crashed...")
				return
			}

			if !strings.Contains(err.Error(), "no such file") {
				t.Errorf("Expected '%s' to contain 'no such file'", err.Error())
			}
		})

		t.Run("Fails on wrong input data", func(t *testing.T) {
			if _, err := OpenConfig("../README.md"); err == nil {
				t.Errorf("Using the README.md as a config file does not crash.")
			}
		})

		t.Run("Can open and parse existing file", func(t *testing.T) {
			var dep1 = Dependency{
				Target: "Foo",
				Source: "Bar",
				Type:   "zip",
			}
			var cfg = Config{Dependencies: []*Dependency{
				&dep1,
			}}
			var fileName = "../test/.edm.test_parse.yaml"

			if err := cfg.WriteFile(fileName); err != nil {
				t.Errorf("Could not write file %s.", fileName)
			}

			newCfg, err := OpenConfig(fileName)

			if err != nil {
				t.Errorf("File %s could not be opened: %s", fileName, err)
			}

			if newCfg.Dependencies[0].Target != cfg.Dependencies[0].Target {
				t.Errorf(
					"File and parsed content do not equal. Expected %s, Got %s",
					cfg.Dependencies[0].Target,
					newCfg.Dependencies[0].Target)
			}
		})
	})
}
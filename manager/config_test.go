package manager

import (
	"gopkg.in/yaml.v2"
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

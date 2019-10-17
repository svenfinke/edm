package manager

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config is providing the structure for the config file
type Config struct {
	Dependencies []*Dependency `yaml:"dependencies"`
}

// OpenConfig is opening und unmarshalling the provided config file
func OpenConfig(filename string) (Config, error) {
	var cfg = Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
			return cfg, err
	}

	return cfg, nil
}

// GenerateFile is generating a yaml file from the config object
func (c *Config) GenerateFile() []byte {
	data, _ := yaml.Marshal(c)

	return data
}

// WriteFile is writing the content from GenerateFile into a file in the filesystem.
func (c *Config) WriteFile(filename string) error {
	return ioutil.WriteFile(filename, c.GenerateFile(), 0644)
}

package manager

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Config is providing the structure for the config file
type Config struct {
	Dependencies []*Dependency `yaml:"dependencies"`
}

// OpenConfig is opening und unmarshalling the provided config file
func OpenConfig(filename string) Config {
	var cfg = Config{}
	if data, err := ioutil.ReadFile(filename); err != nil {
		log.Panic(err)
	} else {
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			log.Panic(err)
		}
	}

	return cfg
}


// GenerateFile is generating a yaml file from the config object
func (c *Config) GenerateFile() []byte{
	if data, err := yaml.Marshal(c); err != nil {
		log.Panic(err)
	} else {
		return data
	}

	return []byte{}
}

// WriteFile is writing the content from GenerateFile into a file in the filesystem.
func (c *Config) WriteFile(filename string) error{
	return ioutil.WriteFile(filename, c.GenerateFile(), 0644)
}
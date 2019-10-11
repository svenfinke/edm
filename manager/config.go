package manager

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Dependencies []*Dependency `yaml:"dependencies"`
}

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

func (c *Config) GenerateFile() []byte{
	if data, err := yaml.Marshal(c); err != nil {
		log.Panic(err)
	} else {
		return data
	}

	return []byte{}
}

func (c *Config) WriteFile(filename string, force bool) error{
	if _, err := os.Open(filename); err == nil {
		if force {
			if err := os.Remove(filename); err != nil {
				return err
			}
		}
	}

	return ioutil.WriteFile(filename, c.GenerateFile(), 0644)
}
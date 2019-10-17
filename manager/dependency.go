package manager

import (
	"errors"
	"fmt"
)

// Dependency is defining a single dependency that must be fetched for the project
type Dependency struct {
	Target string `yaml:"target"`
	Source string `yaml:"source"`
	Type   string `yaml:"type"`
}

// DependencyTypeInterface is used to generate types that can be used in edm
type DependencyTypeInterface interface {
	Fetch(dependency *Dependency) error
	GetInfo() string
}

// DependencyTypes is a map with all available Types
var DependencyTypes map[string]DependencyTypeInterface

func init() {
	DependencyTypes = make(map[string]DependencyTypeInterface)
}

// Fetch is using the type of the dependency and using it to call the correct Fetch of the given type
func (d *Dependency) Fetch() error {
	fmt.Printf("INFO: Fetching %s as %s\n", d.Source, d.Type)

	if depType, ok := DependencyTypes[d.Type]; ok {
		return depType.Fetch(d)
	}

	return errors.New("dependencyType not found")
}

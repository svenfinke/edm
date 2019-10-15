package manager

import (
	"fmt"
	"github.com/svenfinke/edm/helper"
)

type defaultType struct{}

func init() {
	DependencyTypes["default"] = defaultType{}
	DependencyTypes["file"] = defaultType{}
}

// Fetch is downloading the file
func (d defaultType) Fetch(dependency *Dependency) error {
	if err := helper.DownloadFile(dependency.Source, dependency.Target); err != nil {
		return err
	}

	fmt.Printf("INFO: File downloaded to %s\n", dependency.Target)

	return nil
}

// GetInfo is returning a description of the type
func (d defaultType) GetInfo() string {
	return "Downloading the file via HTTP and putting it into the given path."
}

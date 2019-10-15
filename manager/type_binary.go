package manager

import (
	"fmt"
	"github.com/svenfinke/edm/helper"
	"os"
)

type binaryType struct{}

func init() {
	DependencyTypes["binary"] = binaryType{}
}

// Fetch is dowloading given file and making it executable
func (d binaryType) Fetch(dependency *Dependency) error {
	if err := helper.DownloadFile(dependency.Source, dependency.Target); err != nil {
		return err
	}

	fmt.Printf("INFO: File downloaded to %s\n", dependency.Target)

	fileInfo, err := os.Stat(dependency.Target)

	if err != nil {
		return err
	}

	// Change the Filemode
	var newFileMode = fileInfo.Mode() | (1 << 0) | (1 << 3) | (1 << 6)
	if err := os.Chmod(dependency.Target, newFileMode); err != nil {
		return err
	}

	fmt.Printf("INFO: File permissions set from %s to %s\n", fileInfo.Mode(), newFileMode)

	return nil
}

// GetInfo is returning a description of the type
func (d binaryType) GetInfo() string {
	return "Downloading the file via HTTP and putting it into the given path. Additionally making the file executable."
}

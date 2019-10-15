package manager

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/svenfinke/edm/helper"
	"os"
)

const tmpPath = "/tmp/edm"

type zipType struct{}

func init() {
	DependencyTypes["zip"] = zipType{}
}

// Fetch is downloading the file and unzipping it into the target folder
func (z zipType) Fetch(dependency *Dependency) error {
	var tmpFileName = tmpPath + xid.New().String()
	if err := helper.DownloadFile(dependency.Source, tmpFileName); err != nil {
		return err
	}

	fmt.Printf("INFO: File downloaded to /tmp \n")
	files, err := helper.Unzip(tmpFileName, dependency.Target)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: Unzipped %v files to %s", len(files), dependency.Target)

	return os.Remove(tmpFileName)
}

// GetInfo is returning a description of the type
func (z zipType) GetInfo() string {
	return "Downloading the file via HTTP and unzipping it into the given path."
}

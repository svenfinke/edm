package manager

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Dependency is defining a single dependency that must be fetched for the project
type Dependency struct {
	Target string `yaml:"target"`
	Source string `yaml:"source"`
	Type string `yaml:"type"`
}

// Fetch is Downloading the Dependency from Source and Saving it to Target
func (d *Dependency) Fetch () error {
	fmt.Printf("> Fetching '%s', writing to '%s'\n", d.Source, d.Target)

	return downloadFile(d.Source, d.Target)
}

func downloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
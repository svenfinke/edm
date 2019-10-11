package manager

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Dependency struct {
	Target string `yaml:"target"`
	Source string `yaml:"source"`
	Type string `yaml:"type"`
}

func (d *Dependency) Fetch () error {
	fmt.Printf("> Fetching '%s', writing to '%s'\n", d.Source, d.Target)

	return downloadFile(d.Source, d.Target)
}


// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
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
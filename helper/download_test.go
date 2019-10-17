package helper

import (
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	t.Run("Check if invalid target path throws an error", func(t *testing.T) {
		if err := DownloadFile("https://example.com", "../README.md/something"); err == nil {
			t.Errorf("This should have crashed...")
		}
	})

	t.Run("Check if invalid download uri throws an error", func(t *testing.T) {
		var fileName = "../test/check_download"
		if err := DownloadFile("foobar", fileName); err == nil {
			t.Errorf("This should have crashed...")
			os.Remove(fileName)
		}
	})

	t.Run("Check download of sample file", func(t *testing.T) {
		var fileName = "../test/check_download"
		if err := DownloadFile("http://example.com", fileName); err != nil {
			t.Errorf("Download failed: %s", err)
		}
		os.Remove(fileName)
	})
}

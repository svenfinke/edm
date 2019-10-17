package helper

import (
	"os"
	"testing"
)

func TestUnzip(t *testing.T) {
	t.Run("Unzip valid file", func(t *testing.T) {
		if _, err := Unzip("../test/test.zip", "../test/unzip/"); err != nil {
			t.Errorf("Unzipping failed: %s", err)
		}

		os.RemoveAll("../test/unzip")
	})

	t.Run("Check if the function crashes for an invalid archive", func(t *testing.T) {
		if _, err := Unzip("../README.md", "../test/unzip/"); err == nil {
			t.Errorf("Unzipping should've failed")
		}
	})
}
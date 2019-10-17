package manager

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_zipType_Fetch(t *testing.T) {
	t.Run("Checking File Download Error being returned", func(t *testing.T) {
		var dep = Dependency{
			Target: "TestingDownload",
			Source: "WithSomeInvalidPaths",
			Type:   "zip",
		}
		if err := dep.Fetch(); err == nil {
			t.Errorf("Expected an Download Error with given invlaid paths.")
		}
	})

	t.Run("Check if an invalid archive will throw an error", func(t *testing.T) {
		var targetFile = "../test/broken_archive"
		var dep = Dependency{
			Target: targetFile,
			Source: "https://github.com/svenfinke/edm/raw/master/test/fake_binary",
			Type:   "zip",
		}

		if err := dep.Fetch(); err == nil {
			t.Errorf("Fetch should have thrown an error.")
		}
	})

	t.Run("Check if archive is being unzipped and tmpFile is being removed", func(t *testing.T) {
		var targetPath = "../test"
		var dep = Dependency{
			Target: targetPath,
			Source: "https://github.com/svenfinke/edm/raw/master/test/test.zip",
			Type:   "zip",
		}

		files, _ := ioutil.ReadDir("/tmp/edm")
		var numberOfFilesBefore = len(files)

		if err := dep.Fetch(); err != nil {
			t.Errorf("Fetch failed: %s", err)
		}

		files, _ = ioutil.ReadDir("/tmp/edm")
		var numberOfFilesAfter = len(files)

		if _, err := os.Stat(targetPath + "/bar.txt"); err != nil {
			t.Errorf("File bar.txt not found in %s", targetPath)
		}
		if _, err := os.Stat(targetPath + "/foo.txt"); err != nil {
			t.Errorf("File foo.txt not found in %s", targetPath)
		}

		if numberOfFilesBefore < numberOfFilesAfter {
			t.Errorf("The tmp file has not been deleted")
		}
	})
}

func Test_zipType_GetInfo(t *testing.T) {
	t.Run("Zip type is registered and returns the correct string", func(t *testing.T) {
		if depType, ok := DependencyTypes["zip"]; ok {
			var expectedString = zipType{}.GetInfo()
			if !strings.Contains(depType.GetInfo(), expectedString) {
				t.Errorf(
					"Expected '%s' to contain '%s'",
					depType.GetInfo(),
					expectedString,
				)
			}
		} else {
			t.Errorf("Zip DependencyType was not registered.")
		}
	})
}
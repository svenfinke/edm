package manager

import (
	"os"
	"strings"
	"testing"
)

func Test_binaryType_Fetch(t *testing.T) {
	t.Run("Checking File Download Error being returned", func(t *testing.T) {
		var dep = Dependency{
			Target: "TestingDownload",
			Source: "WithSomeInvalidPaths",
			Type:   "binary",
		}
		if err := dep.Fetch(); err == nil {
			t.Errorf("Expected an Download Error with given invlaid paths.")
		}
	})

	t.Run("Check file being downloaded correctly.", func(t *testing.T) {
		var targetFile = "../test/new_fake_binary"
		var dep = Dependency{
			Target: targetFile,
			Source: "https://github.com/svenfinke/edm/raw/master/test/fake_binary",
			Type:   "binary",
		}

		if err := dep.Fetch(); err != nil {
			t.Errorf("Downloading file failed: %s", err)
		}

		fileInfo, err := os.Stat(targetFile)
		if err != nil {
			t.Errorf("%s", err)
		}

		// Using binary AND to only get the executable flags from the fileMode. All 3 of them should be set. Compare
		// this to os.ModePerm where these flags are definitely set.
		if fileInfo.Mode()&((1<<0)|(1<<3)|(1<<6)) != os.ModePerm&((1<<0)|(1<<3)|(1<<6)) {
			t.Errorf("Mode given: %v", fileInfo.Mode()&((1<<0)|(1<<3)|(1<<6)))
		}
	})
}

func Test_binaryType_GetInfo(t *testing.T) {
	t.Run("Binary type is registered and returns the correct string", func(t *testing.T) {
		if depType, ok := DependencyTypes["binary"]; ok {
			var expectedString = binaryType{}.GetInfo()
			if !strings.Contains(depType.GetInfo(), expectedString) {
				t.Errorf(
					"Expected '%s' to contain '%s'",
					depType.GetInfo(),
					expectedString,
				)
			}
		} else {
			t.Errorf("Binary DependencyType was not registered.")
		}
	})
}

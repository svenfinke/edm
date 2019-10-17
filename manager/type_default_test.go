package manager

import (
	"strings"
	"testing"
)

func Test_defaultType_Fetch(t *testing.T) {
	t.Run("Checking File Download Error being returned", func(t *testing.T) {
		var dep = Dependency{
			Target: "TestingDownload",
			Source: "WithSomeInvalidPaths",
			Type:   "default",
		}
		if err := dep.Fetch(); err == nil {
			t.Errorf("Expected an Download Error with given invlaid paths.")
		}
	})

	t.Run("Downloading a valid file should not fail", func(t *testing.T) {
		var targetFile = "../test/broken_archive"
		var dep = Dependency{
			Target: targetFile,
			Source: "https://github.com/svenfinke/edm/raw/master/test/fake_binary",
			Type:   "default",
		}

		if err := dep.Fetch(); err != nil {
			t.Errorf("Fetch should not throw an error.")
		}
	})
}

func Test_defaultType_GetInfo(t *testing.T) {
	t.Run("Default type is registered and returns the correct string", func(t *testing.T) {
		if depType, ok := DependencyTypes["default"]; ok {
			var expectedString = defaultType{}.GetInfo()
			if !strings.Contains(depType.GetInfo(), expectedString) {
				t.Errorf(
					"Expected '%s' to contain '%s'",
					depType.GetInfo(),
					expectedString,
				)
			}
		} else {
			t.Errorf("Default DependencyType was not registered.")
		}
	})

	t.Run("Default type is registered with the alias file and returns the correct string", func(t *testing.T) {
		if depType, ok := DependencyTypes["file"]; ok {
			var expectedString = defaultType{}.GetInfo()
			if !strings.Contains(depType.GetInfo(), expectedString) {
				t.Errorf(
					"Expected '%s' to contain '%s'",
					depType.GetInfo(),
					expectedString,
				)
			}
		} else {
			t.Errorf("Default DependencyType was not registered as file.")
		}
	})
}
package manager

import "testing"

var calledFetch = false

type testType struct{}

func (tt testType) Fetch(dependency *Dependency) error {
	calledFetch = true
	return nil
}
func (tt testType) GetInfo() string {
	return "SomethingAwesome"
}

func TestDependency_Fetch(t *testing.T) {
	t.Run("Run Fetch function with existing type", func(t *testing.T) {
		var dep = Dependency{
			Target: "test",
			Source: "thingy",
			Type:   "testCase",
		}
		DependencyTypes["testCase"] = testType{}
		calledFetch = false
		if err := dep.Fetch(); err != nil {
			t.Errorf("Calling Fetch failed: %s", err)
		}

		if !calledFetch {
			t.Errorf("Fetch has not been called.")
		}
	})

	t.Run("Run Fetch function with non existing type", func(t *testing.T) {
		var dep = Dependency{
			Target: "Foo",
			Source: "Bar",
			Type:   "test",
		}

		if dep.Fetch() == nil {
			t.Errorf("Fetch should have thrown an error.")
		}
	})
}

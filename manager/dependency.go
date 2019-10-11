package manager

type Dependency struct {
	Target string `yaml:"target"`
	Source string `yaml:"source"`
	Type string `yaml:"type"`
}
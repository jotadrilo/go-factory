package config

type Package struct {
	Name           string   `mapstructure:"name"`
	FactoryFileTpl string   `mapstructure:"factory_file_tpl"`
	Include        []string `mapstructure:"include"`
	Exclude        []string `mapstructure:"exclude"`
}

var DefaultFactoryFileTpl = "{{ .ProjectDir }}/{{ .PackageDirRel }}/{{ .Filename }}_factory.go"

type Config struct {
	// Packages are meant to be used within root project runs
	Packages []Package `mapstructure:"packages"`

	// Name is meant to be used with go:generate annotations above go struct types
	// Example:
	//
	// //go:generate go-factory -n Foo
	// type Foo struct {}
	Name string `mapstructure:"name"`

	// File is meant to be used within Name to state the file where the struct is located
	File string `mapstructure:"file"`

	// Project directory. Auto-detected
	ProjectDir string
	// Tool version. Auto-detected
	Version string
}

func NewConfig(version string, dir string) Config {
	return Config{
		ProjectDir: dir,
		Version:    version,
	}
}

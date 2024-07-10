package generator

import (
	"github.com/jotadrilo/go-factory/pkg/config"
	"path/filepath"
)

type Factories struct {
	Config     *config.Config
	Discoverer Discoverer
	Renderer   TreeRenderer
	Writer     DataWriter
}

func NewFactories(cfg *config.Config) *Factories {
	return &Factories{
		Config:     cfg,
		Discoverer: NewFileTreeDiscoverer(cfg),
		Renderer:   NewRenderer(),
		Writer:     NewFileWriter(),
	}
}

// Generate generates the factory code for all the packages specified in the configuration.
func (x *Factories) Generate() error {
	projectDir, err := config.LocateProjectRootDir()
	if err != nil {
		return err
	}

	x.Config.ProjectDir = projectDir

	// There are two run modes:
	// 1. Generate factories from inline annotations
	// 2. Generate factories from configuration
	//
	// If the configuration File and Name fields are set, we are in the first case
	// Otherwise, we are in the second case

	if x.Config.Name != "" && x.Config.File != "" {
		return x.generateFromAnnotation()
	}

	return x.generateFromConfig()
}

func (x *Factories) generateFromConfig() error {
	var pkgTrees []*PackageTree

	for _, pkg := range x.Config.Packages {
		pt, err := x.Discoverer.LoadPackage(pkg)
		if err != nil {
			return err
		}

		pkgTrees = append(pkgTrees, pt)
	}

	for _, pt := range pkgTrees {
		var dataByFile = x.Renderer.FromPackageTree(pt)
		for path, data := range dataByFile {
			if err := x.Writer.Write(path, data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (x *Factories) generateFromAnnotation() error {
	var pkg = config.Package{
		// We don't know the Go package name in this use case (ex. github.com/jotadrilo/go-factory)
		Name:           "",
		Include:        []string{x.Config.Name},
		FactoryFileTpl: config.DefaultFactoryFileTpl,
	}

	pkgFile, err := filepath.Abs(x.Config.File)
	if err != nil {
		return err
	}

	pkgDir := filepath.Dir(pkgFile)

	ft, err := x.Discoverer.LoadFile(pkg, pkgDir, pkgFile)
	if err != nil {
		return err
	}

	var dataByFile = x.Renderer.FromFileTree(ft)

	for path, data := range dataByFile {
		if err := x.Writer.Write(path, data); err != nil {
			return err
		}
	}

	return nil
}

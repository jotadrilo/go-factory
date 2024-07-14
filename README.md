# go-factory

A factory code autogenator for Go

## Installation

### Linux

```
mkdir -p go-factory && \
curl -sSL https://github.com/jotadrilo/go-factory/releases/download/1.0.2/go-factory_1.0.2_linux_x86_64.tar.gz | tar xzf - -C go-factory && \
sudo cp go-factory/go-factory /usr/local/bin/ && \
sudo chmod +x /usr/local/bin/go-factory
```

### MacOS

Homebrew users:

```
brew tap jotadrilo/tap
brew install jotadrilo/tap/go-factory
```

Alternative:

```
mkdir -p go-factory && \
curl -sSL https://github.com/jotadrilo/go-factory/releases/download/1.0.2/go-factory_1.0.2_darwin_arm64.tar.gz | tar xzf - -C go-factory && \
sudo cp go-factory/go-factory /usr/local/bin/ && \
sudo chmod +x /usr/local/bin/go-factory
```

## Configuration

You can configure go-factory by adding annotations to your Go files,
or placing a configuration file in your project root directory.

### Annotations

```golang
package examples

//go:generate go-factory -n A
type A struct {
	Bool   bool
	String string
}

//go:generate go-factory -n B
type B struct {
	Name string
}
```

### Configuration File

The configuration file must include a configuration section per package.

_External packages are not supported._

```yaml
packages:
  - name: 'github.com/jotadrilo/go-factory/examples'
  - name: 'github.com/jotadrilo/go-factory/examples/inner'
    factory_file_tpl: '{{ .ProjectDir }}/{{ .PackageDirRel }}/{{ .Filename }}_factory.go'
    include:
      - A
      - B
    exclude:
      - C
      - D
```

#### Parameters

| Parameter                      | Description                                                                                                                                        |
|--------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| `packages`                     | Array of packages configuration                                                                                                                    |
| `packages[n].name`             | Name of the package to generate factories for                                                                                                      |
| `packages[n].factory_file_tpl` | Go template to configure the factory file output location and name. Default: `'{{ .ProjectDir }}/{{ .PackageDirRel }}/{{ .Filename }}_factory.go'` |
| `packages[n].include`          | Array of struct names to include in the generation. Default: `[]`                                                                                  |
| `packages[n].exclude`          | Array of struct names to exclude from the generation. Default: `[]`                                                                                |  

If both `include` and `exclude` parameters are blank, it will generate factories for all structs in the package.

#### Factory File Template

Let's use the example `examples/inner/foo.go` in this project, assuming that I have cloned the project
in `/home/user/go-factory`:

This template can be configured with the following parameters:

| Parameter       | Description                                                  | Example                 |
|-----------------|--------------------------------------------------------------|-------------------------|
| `ProjectDir`    | Project directory                                            | `/home/user/go-factory` |
| `PackageName`   | Name of the package                                          | `inner`                 |
| `PackageDirRel` | Relative path to the directory of the package                | `examples/inner`        |
| `TypeName`      | Name of the struct type                                      | For `Foo` struct: `Foo` |
| `Filename`      | Name of the file declaring the struct, without its extension | For `Foo` struct: `foo` |

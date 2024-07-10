package generator

import (
	"github.com/jotadrilo/go-factory/pkg/config"
	"go/ast"
)

// Import represents an import declaration
type Import struct {
	Alias string
	Name  string
	Path  string
}

type Imports []*Import

func (x Imports) FindImport(name string) *Import {
	for _, imp := range x {
		if imp.Alias == name || imp.Name == name {
			return imp
		}
	}
	return nil
}

// FileTree represents the required information in a file
type FileTree struct {
	Structs []*Struct
	Imports Imports
}

func (x *FileTree) GetStructs() []*Struct {
	return x.Structs
}

// PackageTree represents the required information in a package
type PackageTree struct {
	Config    config.Package
	Dir       string
	FileTrees []*FileTree
}

func (x *PackageTree) GetStructs() []*Struct {
	var structs []*Struct
	for _, ft := range x.FileTrees {
		structs = append(structs, ft.Structs...)
	}
	return structs
}

// Struct represents a struct in a file
type Struct struct {
	// Root directory of the project
	ProjectDir string

	// Root directory of the package
	PackageRoot string

	// Name of the package
	PackageName string

	// Path to the directory of the package
	PackageDir string

	// Relative path to the directory of the package
	PackageDirRel string

	// Path to the file declaring the struct
	FilePath string

	// Name of the file declaring the struct, without its extension
	Filename string

	// Template to render the output factory file
	FactoryFileTpl string

	// Struct type name
	TypeName string

	// Fields of the struct
	Fields []Field
}

// Field represents a field in the struct with its name and type.
type Field struct {
	Name   string
	Expr   ast.Expr
	Type   string
	Import string
}

func NewField(name string, expr ast.Expr) Field {
	return Field{
		Name:   name,
		Expr:   expr,
		Type:   getFieldNameFromExpr(expr),
		Import: getImportFromExpr(expr),
	}
}

func (x *Field) IsArray() bool {
	_, ok := x.Expr.(*ast.ArrayType)
	return ok
}

func getFieldNameFromExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		elemType := getFieldNameFromExpr(t.Elt)
		return "[]" + elemType
	case *ast.StarExpr:
		elemType := getFieldNameFromExpr(t.X)
		return "*" + elemType
	case *ast.SelectorExpr:
		pkg := getFieldNameFromExpr(t.X)
		return pkg + "." + t.Sel.Name
	case *ast.MapType:
		keyType := getFieldNameFromExpr(t.Key)
		valueType := getFieldNameFromExpr(t.Value)
		return "map[" + keyType + "]" + valueType
	case *ast.FuncType:
		return "func"
	}
	return "unknown"
}

func getImportFromExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.SelectorExpr:
		ident, ok := t.X.(*ast.Ident)
		if !ok {
			return ""
		}
		return ident.Name
	case *ast.StarExpr:
		return getImportFromExpr(t.X)
	case *ast.ArrayType:
		return getImportFromExpr(t.Elt)
	default:
		return ""
	}
}

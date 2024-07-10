package generator

import (
	"log"
)

type TreeRenderer interface {
	FromPackageTree(tree *PackageTree) map[string][]byte
	FromFileTree(tree *FileTree) map[string][]byte
}

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (x *Renderer) FromPackageTree(tree *PackageTree) map[string][]byte {
	var (
		imports Imports
		strcts  []*Struct
	)

	for _, ft := range tree.FileTrees {
		strcts = append(strcts, ft.Structs...)
		imports = append(imports, ft.Imports...)
	}

	return x.renderDataByFile(imports, strcts)
}

func (x *Renderer) FromFileTree(tree *FileTree) map[string][]byte {
	return x.renderDataByFile(tree.Imports, tree.Structs)
}

func (x *Renderer) renderDataByFile(imports Imports, strcts []*Struct) map[string][]byte {
	var (
		dataByFile    = make(map[string][]byte)
		importsByFile = make(map[string][]*Import)
	)

	for _, strct := range strcts {
		for _, field := range strct.Fields {
			if field.Import == "" {
				continue
			}

			imp := imports.FindImport(field.Import)
			if imp != nil {
				importsByFile[strct.FactoryFileTpl] = append(importsByFile[strct.FactoryFileTpl], imp)
				continue
			}

			log.Printf("Unable to locate import %q", field.Import)
		}
	}

	for ix, strct := range strcts {
		log.Printf("Generating factories for %s struct into %s", strct.TypeName, strct.FactoryFileTpl)

		// Add header the first time we are adding the code
		if _, ok := dataByFile[strct.FactoryFileTpl]; !ok {
			dataByFile[strct.FactoryFileTpl] = append(dataByFile[strct.FactoryFileTpl],
				[]byte(generateFactoryCodeHeader(strct))...,
			)

			dataByFile[strct.FactoryFileTpl] = append(dataByFile[strct.FactoryFileTpl],
				[]byte(generateFactoryCodeImports(importsByFile[strct.FactoryFileTpl]))...,
			)
		}

		dataByFile[strct.FactoryFileTpl] = append(dataByFile[strct.FactoryFileTpl],
			[]byte(generateFactoryCode(strct))...,
		)

		if ix < len(strcts)-1 {
			dataByFile[strct.FactoryFileTpl] = append(dataByFile[strct.FactoryFileTpl],
				[]byte("\n\n")...,
			)
		} else if ix == len(strcts) {
			dataByFile[strct.FactoryFileTpl] = append(dataByFile[strct.FactoryFileTpl],
				[]byte("\n")...,
			)
		}
	}

	return dataByFile
}

package generator

import (
	"fmt"
	"strings"
)

func generateFactoryCodeImports(imports []*Import) string {
	if len(imports) == 0 {
		return ""
	}

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("import (\n"))

	var imported = make(map[string]any)
	for _, imp := range imports {
		if imp.Alias != "" {
			_, ok := imported[imp.Alias]
			if ok {
				continue
			}
			imported[imp.Alias] = struct{}{}
			builder.WriteString(fmt.Sprintf("\t%s %q\n", imp.Alias, imp.Path))
		} else {
			_, ok := imported[imp.Path]
			if ok {
				continue
			}
			imported[imp.Path] = struct{}{}
			builder.WriteString(fmt.Sprintf("\t%q\n", imp.Path))
		}
	}

	builder.WriteString(fmt.Sprintf(")\n"))
	builder.WriteString(fmt.Sprintf("\n"))

	return builder.String()
}

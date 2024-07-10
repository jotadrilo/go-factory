package generator

import (
	"fmt"
	"strings"
)

func generateDeepCopyFunction(typeName string, fields []Field) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("func deepCopy%s(src *%s) *%s {\n", typeName, typeName, typeName))
	builder.WriteString(fmt.Sprintf("\tif src == nil {\n"))
	builder.WriteString(fmt.Sprintf("\t\treturn nil\n"))
	builder.WriteString(fmt.Sprintf("\t}\n"))
	builder.WriteString(fmt.Sprintf("\tcopyInstance := *src\n"))

	for _, field := range fields {
		if field.IsArray() {
			builder.WriteString(fmt.Sprintf("\tif src.%s != nil {\n", field.Name))
			builder.WriteString(fmt.Sprintf("\t\tcopyInstance.%s = make(%s, len(src.%s))\n", field.Name, field.Type, field.Name))
			builder.WriteString(fmt.Sprintf("\t\tcopy(copyInstance.%s, src.%s)\n", field.Name, field.Name))
			builder.WriteString(fmt.Sprintf("\t}\n"))
		} else if strings.HasPrefix(field.Type, "*") {
			builder.WriteString(fmt.Sprintf("\tif src.%s != nil {\n", field.Name))
			builder.WriteString(fmt.Sprintf("\t\tcopyInstance.%s = new(%s)\n", field.Name, field.Type[1:]))
			builder.WriteString(fmt.Sprintf("\t\t*copyInstance.%s = *src.%s\n", field.Name, field.Name))
			builder.WriteString(fmt.Sprintf("\t}\n"))
		}
	}

	builder.WriteString(fmt.Sprintf("\treturn &copyInstance\n"))
	builder.WriteString("}\n")

	return builder.String()
}

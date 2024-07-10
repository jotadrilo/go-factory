package generator

import (
	"fmt"
	"strings"
)

func generateFactoryCode(strct *Struct) string {
	var (
		builder  strings.Builder
		typeName = strct.TypeName
		fields   = strct.Fields
	)

	builder.WriteString(fmt.Sprintf("// Factory%s is a helper factory to ease creating data of type %s\n", typeName, typeName))
	builder.WriteString(fmt.Sprintf("type Factory%s struct {\n", typeName))
	builder.WriteString(fmt.Sprintf("\tFactory %s\n", typeName))
	builder.WriteString(fmt.Sprintf("}\n"))
	builder.WriteString(fmt.Sprintf("\n"))
	builder.WriteString(fmt.Sprintf("func NewFactory%s() *Factory%s {\n", typeName, typeName))
	builder.WriteString(fmt.Sprintf("\treturn &Factory%s{}\n", typeName))
	builder.WriteString(fmt.Sprintf("}\n"))
	builder.WriteString(fmt.Sprintf("\n"))

	for _, field := range fields {
		if field.IsArray() {
			// WithField method
			builder.WriteString(fmt.Sprintf("func (f *Factory%s) With%s(values ...%s) *Factory%s {\n", typeName, field.Name, field.Type[2:], typeName))
			builder.WriteString(fmt.Sprintf("\tf.Factory.%s = values\n", field.Name))
			builder.WriteString(fmt.Sprintf("\treturn f\n"))
			builder.WriteString(fmt.Sprintf("}\n"))
			builder.WriteString(fmt.Sprintf("\n"))

			// AddField method
			builder.WriteString(fmt.Sprintf("func (f *Factory%s) Add%s(values ...%s) *Factory%s {\n", typeName, field.Name, field.Type[2:], typeName))
			builder.WriteString(fmt.Sprintf("\tf.Factory.%s = append(f.Factory.%s, values...)\n", field.Name, field.Name))
			builder.WriteString(fmt.Sprintf("\treturn f\n"))
			builder.WriteString(fmt.Sprintf("}\n"))
			builder.WriteString(fmt.Sprintf("\n"))
		} else {
			builder.WriteString(fmt.Sprintf("func (f *Factory%s) With%s(value %s) *Factory%s {\n", typeName, field.Name, field.Type, typeName))
			builder.WriteString(fmt.Sprintf("\tf.Factory.%s = value\n", field.Name))
			builder.WriteString(fmt.Sprintf("\treturn f\n"))
			builder.WriteString(fmt.Sprintf("}\n"))
			builder.WriteString(fmt.Sprintf("\n"))
		}
	}

	builder.WriteString(fmt.Sprintf("func (f *Factory%s) Build() *%s {\n", typeName, typeName))
	builder.WriteString(fmt.Sprintf("\treturn deepCopy%s(&f.Factory)\n", typeName))
	builder.WriteString(fmt.Sprintf("}\n"))
	builder.WriteString(fmt.Sprintf("\n"))

	builder.WriteString(generateDeepCopyFunction(typeName, fields))

	return strings.TrimSpace(builder.String())
}

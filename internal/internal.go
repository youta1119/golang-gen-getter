package internal

import (
	"github.com/youta1119/golang-gen-getter/internal/generator"
	"go/types"
)

func GenerateGetter(typeName string, sourceDirs []string) ([]byte, error) {
	pkg, err := parsePackage(sourceDirs)
	if err != nil {
		return nil, err
	}

	files, err := parseFiles(pkg.GoFiles)
	if err != nil {
		return nil, err
	}

	st, err := findTargetStructType(typeName, files)
	if err != nil {
		return nil, err
	}

	fields := make([]*generator.Field, 0, len(st.Fields.List))
	for _, field := range st.Fields.List {
		fields = append(fields, &generator.Field{
			Name:     field.Names[0].Name,
			TypeName: types.ExprString(field.Type),
		})
	}
	gen := generator.GetterGenerator{
		PkgName:  pkg.Name,
		TypeName: typeName,
		Fields:   fields,
	}
	return gen.Generate()
}

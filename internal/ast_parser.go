package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func parsePackage(dirPattern []string) (*packages.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		Tests: false,
	}, dirPattern...)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("error; %d packages found", len(pkgs))
	}
	return pkgs[0], nil
}

func parseFiles(files []string) ([]*ast.File, error) {
	fset := token.NewFileSet()

	astFiles := make([]*ast.File, len(files))
	for i, file := range files {
		parsed, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file: %w", err)
		}
		astFiles[i] = parsed
	}
	return astFiles, nil
}


func findTargetStructType(typeName string, astFiles []*ast.File) (*ast.StructType, error) {
	for _, astFile := range astFiles {
		for _, decl := range astFile.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structName := typeSpec.Name.Name
				if typeName != structName {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				return structType, nil
			}
		}
	}

	return nil, fmt.Errorf("can't found given typeName [given=%s]", typeName)
}
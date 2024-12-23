package main

import (
	"embed"
	_ "embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"

	"golang.org/x/tools/imports"
)

type TemplateValues struct {
	PackageName string
	EnumName    string
	Values      []string
}

//go:embed tmpl/*
var tmpls embed.FS

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: enumgen constant_type_name input_file.go output_file.go")
		os.Exit(1)
	}
	mainFn(args[0], args[1], args[2])
}

func mainFn(inputFilePath, outputFilePath, outputPackage string) {
	templates, err := template.New("").
		ParseFS(tmpls, "tmpl/*")
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, inputFilePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	templateValues := TemplateValues{
		PackageName: node.Name.Name,
		EnumName:    outputPackage,
		Values:      []string{},
	}

	// list over const values
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.CONST {
			continue
		}

		// NOTE: need to handle multiple enum types
		currentType := ""
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(valueSpec.Names) != 1 {
				continue
			}
			if valueSpec.Type != nil {
				currentType = valueSpec.Type.(*ast.Ident).Name
			}
			if currentType != templateValues.EnumName {
				continue
			}
			templateValues.Values = append(templateValues.Values, valueSpec.Names[0].Name)
		}
	}

	outFile, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = templates.ExecuteTemplate(outFile, "enum.tmpl", templateValues)
	if err != nil {
		panic(err)
	}
	err = outFile.Close()
	if err != nil {
		panic(err)
	}

	outputWithImports, err := imports.Process(outputFilePath, nil, &imports.Options{
		Fragment:   false,
		AllErrors:  true,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outputFilePath, outputWithImports, 0644)
	if err != nil {
		panic(err)
	}
}

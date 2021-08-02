package main

import (
	"fmt"
)

type ModuleDoc struct {
	AbsolutePath  string
	ImportPath    string
	Packages      []Package
	SimpleExports SimpleExportsByType
}

type ExportType string
type SimpleExportsByType map[ExportType][]string

func (m ModuleDoc) Print() {
	fmt.Println("ModuleDoc{ ", "AbsolutePath =", m.AbsolutePath, " ImportPath =", m.ImportPath)
	for exportType, exports := range m.SimpleExports {
		fmt.Println(" - ", exportType, ":", exports)
	}
}

type Snippet interface {
	Snip() string
}

type Function struct {}
type Method struct {}
type Typedef struct{}

type CodeDefinition struct {
	Functions []Function
	Methods []Method
	Typedefs []Typedef
}

type PackageFile struct {
	Pkg *Package
	Code CodeDefinition
}

type Package struct {
	Name string
	Path string
	ExportedCode CodeDefinition
	Files []PackageFile
}

package lib

import (
	"go/token"
	"fmt"
	"go/parser"
	"log"
	"go/ast"
	"original/go/src/path/filepath"
)

func FoundHandlablePackages(set *token.FileSet, baseDir string, pkgPathList []string) (pkgList []*PkgCreator)  {
	counter := 0
	for _, path := range pkgPathList {
		pkgCreator := new(PkgCreator)
		pkgCreator.Path = path
		pkgCreator.BaseDir = baseDir

		pkgCreator.Init(set)
		if pkgCreator.IsValid() {
			pkgList = append(pkgList, pkgCreator)
			fmt.Printf("%s: %s\n", path, pkgCreator.Name)
			counter++
			if counter > 10 {
				break
			}
		}
	}
	return
}

func LoadTemplateFunc(set *token.FileSet, baseDir string) []ast.Stmt {
	node, err := parser.ParseFile(set, filepath.Join(baseDir, "templates/template-func.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var templateIf *ast.FuncDecl

	ast.Inspect(node, func(n ast.Node) bool {
		// handle function declarations without documentation
		ifSample, ok := n.(*ast.FuncDecl)
		if ok {
			templateIf = ifSample
			return false
		}
		return true
	})
	return templateIf.Body.List
}

func loadTemplateStruct(set *token.FileSet, baseDir string) *ast.File {
	node, err := parser.ParseFile(set, filepath.Join(baseDir, "templates/template-struct.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return node
}
func loadTemplateMirror(set *token.FileSet, baseDir string) *ast.File {
	node, err := parser.ParseFile(set, filepath.Join(baseDir, "templates/template-mirror.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return node
}
func LoadTemplateMock(set *token.FileSet, baseDir string) *ast.File {
	node, err := parser.ParseFile(set, filepath.Join(baseDir, "templates/template-struct.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return node
}
//func loadTemplateStruct(set *token.FileSet, baseDir string) *ast.TypeSpec {
//	node, err := parser.ParseFile(set, filepath.Join(baseDir, "templates/template-struct.go"), nil, parser.ParseComments)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var templateType *ast.TypeSpec
//	ast.Inspect(node, func(n ast.Node) bool {
//		// handle function declarations without documentation
//		strSample, ok := n.(*ast.TypeSpec)
//		if ok {
//			templateType = strSample
//			return false
//		}
//		return true
//	})
//	return templateType
//}
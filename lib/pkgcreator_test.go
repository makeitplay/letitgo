package lib

import (
	"testing"
	"os"
	"runtime"
	"path/filepath"
	"path"
	"go/token"
	"go/parser"
	"go/ast"
	"go/printer"
	"log"
	"original/go/src/fmt"
	"go/format"
)
var fs = token.NewFileSet()
var pkgCreator *PkgCreator
var baseDir string
func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	baseDir, _ = filepath.Abs(path.Dir(filename) + "/../")
	pkgCreator = new(PkgCreator)
	pkgCreator.Path = "testdata/"
	pkgCreator.BaseDir = baseDir
	fmt.Printf("%s \n", baseDir)
	pkgCreator.Init(fs)
	//pkgCreator.MockUp()
	os.Exit(m.Run())
}

func TestPkgCreator_CustomizeCallback(t *testing.T) {
	//parser.ParseFile(fs, filepath.Join(baseDir, "templates/template-func.go"), nil, parser.Trace)

	file, err := parser.ParseFile(fs, filepath.Join(baseDir, "lib/testdata/inputLib.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		templateCond := LoadTemplateFunc(fs, baseDir)
		// handle function declarations without documentation
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			fmt.Printf("\n\n##### one more ##### \n")
			pkgCreator.CustomizeCallback(fn, templateCond)
		}
		return true
	})

	f2, _ := os.Create("output.nogo")

	defer f2.Close()

	if err := printer.Fprint(f2, pkgCreator.fs,file); err != nil {
		log.Fatal(err)
	}
	//parser.ParseFile(fs, "output.nogo", nil, parser.Trace)
}

func TestPkgCreator_writeFuncInMirror(t *testing.T) {
	pkgCreator.pkgMirror = loadTemplateMirror(fs, baseDir)
	file, err := parser.ParseFile(fs, filepath.Join(baseDir, "lib/testdata/inputLib.go"), nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var outputTest *ast.File
	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			fmt.Printf("\n\n##### one more ##### \n")
			outputTest = pkgCreator.writeFuncInMirror(file, fn)
		}
		return true
	})

	f2, _ := os.Create("output.nogo")

	format.Node(f2, pkgCreator.fs, outputTest)

	defer f2.Close()

	//if err := printer.Fprint(f2, pkgCreator.fs,outputTest); err != nil {
	//	log.Fatal(err)
	//}
}
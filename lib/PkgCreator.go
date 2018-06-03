package lib

import (
	"go/token"
	"go/parser"
	"go/ast"
	"regexp"
	"fmt"
	"os"
	"go/printer"
	"log"
	"path/filepath"
	"strings"
)

type PkgCreator struct {
	Path string
	Name string
	BaseDir string
	OriginalPkg *ast.Package
	fs *token.FileSet
	fakeStruct  *ast.File
}

func (c *PkgCreator) Init(fs *token.FileSet) error {
	pkgs, err := parser.ParseDir(fs, c.Path, nil, 0)

	if err != nil {
		return err
	}
	for pkgName, pkg := range pkgs {
		ast.PackageExports(pkg)
		c.Name = pkgName
		c.OriginalPkg = pkg
		c.fs = fs
		//upgradePkgFiles(pkg)
		//fmt.Printf("--- Pkg %s%s ----\n", lib,pkgName)
	}
	return nil
}
func (c *PkgCreator) IsValid() bool {
	patternNegative := regexp.MustCompile(`(_test|builtin)`)
	return c.Name != "" && !patternNegative.MatchString(c.Name)
}
func (c *PkgCreator) MockUp() error {
	//totalFiles := len(c.OriginalPkg.Files)
	//totalParsed := 0
	//c.fakeStruct = loadTemplateStruct(c.fs, c.BaseDir)
	//c.fakeStruct.Imports = []*ast.ImportSpec{}
	for fileName := range c.OriginalPkg.Files {
		patternNegative := regexp.MustCompile(`_test\.go`)
		if !patternNegative.MatchString(fileName) {
			//fmt.Printf("%s\n", fileName)
			node, err := parser.ParseFile(c.fs, fileName, nil, parser.ParseComments)
			if err != nil {
				return err
			}
			c.createStruct(node)
			c.SaveStruct(node)
			//WriteIfStm(node, fileName)

			//totalParsed++
			//fmt.Printf("Total %d %%\n", totalParsed/totalFiles)
		}
	}
	return nil
}
func (c *PkgCreator) createStruct(file *ast.File) {
	templateIf := LoadTemplateFunc(c.fs, c.BaseDir)
	c.fakeStruct = loadTemplateStruct(c.fs, c.BaseDir)
	//fmt.Printf("exported function %s")
	//sair := false

	//ast.Inspect(file, func(n ast.Node) bool {
	//	mprt, ok := n.(*ast.ImportSpec)
	//	if ok {
	//		if mprt.Path != nil {
	//			c.fakeStruct.Imports = append(c.fakeStruct.Imports, mprt)
	//
	//			importC := &ast.GenDecl{
	//				TokPos: c.fakeStruct.Package,
	//				Tok:    token.IMPORT,
	//				Specs:  []ast.Spec{&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value:  mprt.Path.Value}}},
	//			}
	//			c.fakeStruct.Decls = append([]ast.Decl{importC}, c.fakeStruct.Decls...)
	//
	//
	//			fmt.Printf("IMPORTING: %s\n", mprt.Path.Value)
	//		}
	//	}
	//	return true
	//})
	foundExported := false
	ast.Inspect(file, func(n ast.Node) bool {
		// handle function declarations without documentation
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if fn.Name.IsExported() && fn.Recv == nil {
				foundExported = true

				//line := c.fs.Position(fn.Pos()).Line
				//filenm := fn.Name.Name
				//declaretion := readStatement(fileName, fn.Pos(), fn.End())
				//fmt.Printf("exported function %s" +
				//	"found on line %d: \n\t%s\n", declaretion, line, filenm)

				fld := &ast.Field{}
				fld.Type = fn.Type
				fld.Names = []*ast.Ident{}
				fld.Names = append(fld.Names, fn.Name)

				c.writeFieldFakeStrc(fld)

				//sair = true
				fn.Body.List = append([]ast.Stmt{templateIf}, fn.Body.List...)
			}
		}
		return true
	})
	if foundExported {
		c.includeImport(file)
	}
}

func (c *PkgCreator) writeFieldFakeStrc(field *ast.Field) {
	ast.Inspect(c.fakeStruct, func(n ast.Node) bool {
		strSample, ok := n.(*ast.StructType)
		if ok {
			strSample.Fields.List = append(strSample.Fields.List, field)
		}
		return true
	})
}
func (c *PkgCreator) SaveStruct(file *ast.File) {
	ast.Inspect(c.fakeStruct, func(n ast.Node) bool {
		strSample, ok := n.(*ast.TypeSpec)
		if ok {
			strSample.Name.Name = c.FInalStructName()
			strDelc := &ast.GenDecl{
				TokPos: file.Pos(),
				Tok:    token.TYPE,
				Specs:  []ast.Spec{strSample},
			}
			file.Decls = append(file.Decls, strDelc)
		}
		return true
	})

	f2, _ := os.Create(c.pathToFakeFile())

	defer f2.Close()

	if err := printer.Fprint(f2, c.fs,file); err != nil {
		log.Fatal(err)
	}
}
func (c *PkgCreator) pathToFakeFile() string {
	return fmt.Sprintf(filepath.Join(c.BaseDir, "fake_env/go/src/letgo/%s.lixo"), c.Name)
}
func (c *PkgCreator) FInalStructName() string {
	return fmt.Sprintf("Mck%s", strings.Title(c.Name))
}
func (c *PkgCreator) includeImport(file *ast.File) {

	strDelc := &ast.GenDecl{
		TokPos: file.Package,
		Tok:    token.IMPORT,
		Specs:  []ast.Spec{&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: "letgo"}}},
	}
	file.Decls = append([]ast.Decl{strDelc}, file.Decls...)
}




package main

import (
	"path/filepath"
	"os"
	"fmt"
	"regexp"
	"go/token"
	"go/ast"
	"go/printer"
	"log"
	//"letgo/lib"
	"go/parser"
	"letgo/lib"
)

const base = "/vagrant/go/src/letgo"
const source = base + "/../original/go/src"
var pkgPathList = []string{}

func main()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Pay %v", r)
		}
	}()
	//trainning()
	//os.Exit(0)
	err := filepath.Walk(source, scanDir)
	if err != nil {
		panic(err)
	}
	var fs = token.NewFileSet()
	pkgList := lib.FoundHandlablePackages(fs, base, pkgPathList)
	fmt.Printf("%d Packages found.\n", len(pkgList))
	for _, pkgCreator := range pkgList {
		err = pkgCreator.MockUp()
		if err != nil {
			panic(err)
		}
	}

	return
}
func trainning() {
	var fs = token.NewFileSet()

	templateIf, _ := parser.ParseFile(fs, filepath.Join(base, "templates/template-func.go"), nil, parser.Trace)
	//templateIf := lib.LoadTemplateFunc(fs, base)

	ast.Inspect(templateIf, func(n ast.Node) bool {
		// handle function declarations without documentation

		fmt.Printf("Achou: %T -> %v\n", n, n)
		tipoFun, ok := n.(*ast.BasicLit)
		if ok {
			tipoFun.Value = "\"qualquer label\""
			//Pegando o tipo que conversao da interface
			//fmt.Printf("AchSSSou: %v\n", tipoFun.Value)
			//} else {
			//	returnType, ok2 := n.(*ast.ReturnStmt)
			//	if ok2 {
			//		fmt.Printf("ENTROU2: %v", len(returnType.Results))
			//	}
		}
		//myIf, ok := n.(*ast.FuncType)
		//if ok {
		//
		//	ast.Inspect(myIf.Body.List[0], func(n ast.Node) bool {
		//		tipoFun, ok := n.(*ast.Ident)
		//		if ok {
		//			//Pegando o tipo que conversao da interface
		//			fmt.Printf("Achou: %v\n", tipoFun.Name)
		//		//} else {
		//		//	returnType, ok2 := n.(*ast.ReturnStmt)
		//		//	if ok2 {
		//		//		fmt.Printf("ENTROU2: %v", len(returnType.Results))
		//		//	}
		//		}
		//		return true
		//	})
		//
		//	//fmt.Printf("%v", ifSample.Decl)
		//	return false
		//}
		return true
	})

	f2, _ := os.Create("test.go")

	defer f2.Close()

	if err := printer.Fprint(f2, fs,templateIf); err != nil {
		log.Fatal(err)
	}
}



func scanDir(path string, f os.FileInfo, err error) error {
	patternNegative := regexp.MustCompile(`(_test|testdata)`)
	if f.IsDir() && !patternNegative.MatchString(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		pkgPathList = append(pkgPathList, path)
	}
	return nil
}




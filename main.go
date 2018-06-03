package letgo

import (
	"path/filepath"
	"os"
	"fmt"
	"regexp"
	Llib "letgo/lib"
	"go/token"
	"go/ast"
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
	var fs = token.NewFileSet()
	templateIf := Llib.LoadTemplateFunc(fs, base)

	ast.Inspect(templateIf, func(n ast.Node) bool {
		// handle function declarations without documentation

		myIf, ok := n.(*ast.IfStmt)
		if ok {

			ast.Inspect(myIf.Body.List[0], func(n ast.Node) bool {
				tipoFun, ok := n.(*ast.TypeAssertExpr)
				if ok {
					//Pegando o tipo que conversao da interface
					fmt.Printf("ENTROU: %v", tipoFun.Type)
				} else {
					returnType, ok2 := n.(*ast.ReturnStmt)
					if ok2 {
						fmt.Printf("ENTROU2: %v", len(returnType.Results))
					}
				}
				return true
			})

			//fmt.Printf("%v", ifSample.Decl)
			return false
		}
		return true
	})


	//err := filepath.Walk(source, scanDir)
	//if err != nil {
	//	panic(err)
	//}

	//pkgList := lib.FoundHandlablePackages(fs, base, pkgPathList)
	//fmt.Printf("%d Packages found.\n", len(pkgList))
	//for _, pkgCreator := range pkgList {
	//	err = pkgCreator.MockUp()
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	return
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




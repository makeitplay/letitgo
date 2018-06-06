package testdata

import (
	"original/go/src/strings"
	"github.com/cznic/mathutil"
)

func ANiceFunction(a mathutil.FCBig, testIt bool) {
	if a.Pos() != nil {
		somevar := 2
		somevar ++
		//a comment from ANiceFunction
	}
	return
}

func MyNiceFunction(b string, outro ...interface{}) {
	if strings.Contains(b, "something") {
		somevar := 2

		somevar ++
		//other comment at MyNiceFunction
	}
	return
}

func AnotehrNiceFunction(c string, maisum []map[string]interface{}) int {
	if strings.Contains(c, "something") {
		somevar := 2

		somevar ++
		//funny comment from AnotehrNiceFunction
	}
	outra := []interface{}{"isso", "aquilo"}
	MyNiceFunction(c, outra...)
	return 3
}

func LastNiceFunction(d string, maisumoutro []map[string]interface{}) (cool int, notCool error) {
	if strings.Contains(d, "something") {
		somevar := 2

		somevar ++
		//last comment in LastNiceFunction
	}
	return
}
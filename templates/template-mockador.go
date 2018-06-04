package mockador

import (
	"mymock"
)

type mockador struct {
	mymock mymock.Ju
	Fmt *FmtMirror
}
func (m *mockador)Init()  {
	mymock.Mock().Mocado = true
	mymock.Mock().Mocks =  map[string]interface{}{}


	m.Fmt = new(FmtMirror)
}

func NewMocador() *mockador {
	m := new(mockador)
	m.Init()
	return m
}





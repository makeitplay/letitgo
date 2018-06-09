package mockador

import (
	"mymock"
)

type Mockador struct {
	//mymock mymock.Breakin

	//gerado
	Fmt *FmtMirror
}

func NewMocador() *Mockador {
	m := new(Mockador)
	m.Init()
	return m
}

func (m *Mockador)Init()  {
	mymock.Mock().Mocks =  map[string]interface{}{}

	//gerado
	m.Fmt = new(FmtMirror)
}






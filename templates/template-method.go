package letgo

func (a *FmtMirror) Printf(f func(format string, a ...interface{}) (n int, err error)) {
	mymock.Mock().Mocks["fmt.Printf"] = f
}


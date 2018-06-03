package letgo

func Sample() int {
	if mymock.Mock().IsMockado("fmt.Printf") {
		return mymock.Mock().Fn("fmt.Printf").(MockPrintf)(format, a...)
	}
	return 0
}

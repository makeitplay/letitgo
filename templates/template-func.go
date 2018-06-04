package letgo

func Sample() int {
	letitgoLabel := "fmt.Printf"
	if mymock.Mock(letitgoLabel).IsMockado() {
		return mymock.Mock().Fn(letitgoLabel).(func(args int) (bool))()
	}
}

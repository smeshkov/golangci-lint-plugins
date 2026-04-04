package example

func traditional() {
	for i := 0; i < 10; i++ { // want `use Go 1.22 'for ... := range N' syntax instead of traditional for-loop`
		_ = i
	}
}

func modernRange() {
	for i := range 10 {
		_ = i
	}
}

func noInit() {
	i := 0
	for ; i < 10; i++ {
		_ = i
	}
}

func noCond() {
	for i := 0; ; i++ {
		break
		_ = i
	}
}

func noPost() {
	for i := 0; i < 10; {
		i++
	}
}

func decrement() {
	for i := 10; i > 0; i-- {
		_ = i
	}
}

func lessEqual() {
	for i := 0; i <= 10; i++ {
		_ = i
	}
}

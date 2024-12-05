package utils

func Panic_fn(err error) {
	if err != nil {
		panic(err)
	}
}

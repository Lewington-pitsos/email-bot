package helpers

func CheckSafe(err error) {
	if err != nil {
		panic(err)
	}
}

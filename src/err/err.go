package err

//PanicError is a method to check the error and panic it
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

package iteration

func Repeat(char string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}

func RepeatWithBytesArray(char string, n int) string {
	repeated := make([]byte, n)
	for i := range repeated {
		repeated[i] = char[0]
	}
	return string(repeated)
}

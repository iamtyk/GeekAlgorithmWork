package week8

func toLowerCase(str string) string {
	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			b[i] = str[i] + 32
		}
	}
	return string(b)
}

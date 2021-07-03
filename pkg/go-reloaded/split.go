package reloaded

func Split(s, sep string) []string {
	var sArr = []rune(s)
	var sepArr = []rune(sep)
	var sepLen = len(sep)
	var start = 0
	var output []string

	for i, v := range sArr {
		if v == sepArr[0] && s[i:i+sepLen] == sep {
			output = append(output, s[start:i])
			start = i + sepLen
		}
	}
	output = append(output, s[start:])

	return output
}

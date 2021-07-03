package reloaded

func SplitWhiteSpaces(s string) []string {
	var output []string
	var start = 0

	for i, v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			output = append(output, s[start:i])
			start = i + 1
		}
	}
	output = append(output, s[start:])

	return output
}

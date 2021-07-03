package reloaded

func Atoi(s string) int {
	var sign = s[0]
	var output = 0

	if sign == '+' || sign == '-' {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	for _, v := range s {
		if v >= '0' && v <= '9' {
			v -= 48
			output = output*10 + int(v)
		} else {
			return 0
		}
	}

	if sign == '-' {
		output = -output
	}

	return output
}

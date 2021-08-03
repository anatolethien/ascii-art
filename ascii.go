package ascii

func Ascii(args []string, banner string) {
	var output = ""
	for _, str := range args {
		for i := 0; i < 8; i++ {
			for _, v := range str {
				output += GetLine(banner, 1+int(v-32)*9+i)
			}
			output += "\n"
		}
	}
	return output
}

package ascii

import (
// "github.com/anatolethien/go-reloaded"
)

func Classic() string {
	var args = GetArgs()
	var banner = "standard"
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

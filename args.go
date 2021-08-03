package ascii

import (
	"github.com/anatolethien/go-reloaded"
	"os"
)

func GetArgs() []string {
	var input []string = os.Args[1:]
	return reloaded.Split(reloaded.Join(input, " "), "\\n")
}

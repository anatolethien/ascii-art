package ascii

import (
    "github.com/anatolethien/ascii-art/pkg/go-reloaded"
)

func Classic() []string {
    var args = GetArgs()
	// var file = GetFile("standard")
    var splitInput = reloaded.Split(ArrToStr(args, 0), "\n")
    return splitInput
}

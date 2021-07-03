package ascii

import (
    "os"
)

func GetArgs() []string {
    return os.Args[1:]
}

func ArrToStr(args []string, start int) string {
    var output = ""
    for _, v := range args[start:] {
        output += v
    }
    return output
}

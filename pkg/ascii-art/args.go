package ascii

import (
    "os"
)

func GetArgs() []string {
    return os.Args[1:]
}

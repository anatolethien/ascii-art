package ascii

import (
	"fmt"
	"io/ioutil"
)

func GetFile(banner string) []byte {
	var file, _ = ioutil.ReadFile(fmt.Sprintf("./assets/%s.txt", banner))
	return file
}

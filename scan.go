package ascii

import (
	"bufio"
	"log"
	"os"
)

func GetFile(banner string) *os.File {
	var file, err = os.Open(fmt.Sprintf("./%s.txt", banner))
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func GetLine(banner string, lineNum int) string {
	var file = GetFile(banner)
	var scanner = bufio.NewScanner(file)
	var lineScan = 0
	var line = ""
	for scanner.Scan() {
		if lineScan == lineNum {
			line = scanner.Text()
		}
		lineScan++
	}
	return line
}

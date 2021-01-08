package ascii_art

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func scanner(format string, num int) string {

	var file, err = os.Open("./format/" + format + ".txt")
	var scanner = bufio.NewScanner(file)
	var lineScan = 0
	var line = ""

	if err != nil {

		errorExit()

	}

	for scanner.Scan() {

		if lineScan == num {

			line = scanner.Text()

		}

		lineScan++

	}

	return line
}

func newlineCheck(text string) bool {

	var previous = '0'

	for _, v := range text {

		if v == 'n' && previous == '\\' {
			return true
		}

		previous = v

	}

	return false
}

func formatCheck(args []string) bool {

	switch args[2] {

	case "standard":
		return true

	case "shadow":
		return true

	case "thinkertoy":
		return true

	case "doom":
		return true

	}

	return false
}

func colorFlag(args []string) string {

	for _, v := range args[2:] {

		if len(args[2]) > 8 && v[:8] == "--color=" {

			var color = v[8:]

			if color == "white" || color == "orange" || color == "magenta" || color == "lightblue" || color == "yellow" || color == "lime" || color == "pink" || color == "gray" || color == "lightgray" || color == "cyan" || color == "purple" || color == "blue" || color == "brown" || color == "green" || color == "red" || color == "black" {

				return colorRGB(color)

			} else {

				unknownColor()
				colorUsage()
				errorExit()

			}

		} else {

			incorrectSyntax()
			colorUsage()
			errorExit()

		}

	}
	return "\u001b[0m"
}

func colorRGB(color string) string {

	switch color {

	case "white":
		color = "\u001b[38;2;255;255;255m"

	case "orange":
		color = "\u001b[38;2;255;127;0m"

	case "magenta":
		color = "\u001b[38;2;255;0;255m"

	case "lightblue":
		color = "\u001b[38;2;63;191;255m"

	case "yellow":
		color = "\u001b[38;2;255;255;0m"

	case "lime":
		color = "\u001b[38;2;63;255;0m"

	case "pink":
		color = "\u001b[38;2;255;127;127m"

	case "gray":
		color = "\u001b[38;2;63;63;63m"

	case "lightgray":
		color = "\u001b[38;2;127;127;127m"

	case "cyan":
		color = "\u001b[38;2;0;255;255m"

	case "purple":
		color = "\u001b[38;2;127;0;191m"

	case "blue":
		color = "\u001b[38;2;0;0;255m"

	case "brown":
		color = "\u001b[38;2;127;63;0m"

	case "green":
		color = "\u001b[38;2;0;255;0m"

	case "red":
		color = "\u001b[38;2;255;0;0m"

	case "black":
		color = "\u001b[38;2;0;0;0m"
	}

	return color
}

func outputFlag(args []string) string {

	var output = ""

	for _, v := range args[3:] {

		if len(args[3]) > 9 && v[:9] == "--output=" {

			output = v[9:]

		} else {

			incorrectSyntax()
			outputUsage()
			errorExit()

		}

	}

	return output
}

func justifyFlag(args []string) string {

	var align = ""

	for _, v := range args[3:] {

		if len(args[3]) > 8 && v[:8] == "--align=" {

			align = v[8:]

			if align == "left" || align == "right" || align == "center" || align == "justify" {

				return align
			} else {

				unknownAlign()
				justifyUsage()
				errorExit()

			}

		} else {

			incorrectSyntax()
			justifyUsage()
			errorExit()

		}

	}

	return align
}

func justifyLeft(text string, format string, result string)  {

	for i := 0 ; i < 8 ; i++ {

		for _, v := range text {

			result += scanner(format, 1 + int(v - 32) * 9 + i)

		}

		fmt.Println(result)
		result = ""

	}

}

func justifyJustify(text string, result string, format string, terminalWidth int)  {

	var splitWSText = splitWhiteSpaces(text)
	var array2 = make([][]string, len(splitWSText))
	result = ""
	var j = 0

	for i := 0; i < 8; i++ {

		j = 0

		for _, v := range text {

			if v != 32 {

				result += scanner(format, 1+int(v-32)*9+i)

			} else if v == 32 && result != "" {

				array2[j] = append(array2[j], result)
				result = ""
				j++

			}

		}

		array2[j] = append(array2[j], result)
		result = ""

	}

	var resultLen = 0

	for _, array := range array2 {

		resultLen += len(array[0])

	}

	if len(splitWSText) == 1 {

		justifyLeft(text, format, result)

		return
	}

	var spaces = (terminalWidth - resultLen) / (len(splitWSText) - 1)

	for i := 0; i < 8; i++ {

		for k, v := range array2 {

			fmt.Print(v[i])

			if k != len(array2)-1 {

				fmt.Print(fillSpaces(spaces))

			}

		}

		fmt.Println()

	}

}

func terminalWidth() int {

	//enters a command to retrieve information about the terminal
	var out1, err = exec.Command("tput", "cols").Output() //powershell.exe -noprofile -command $host.ui.rawui.WindowSize.width ou MODE
	//converts the needed information into a string
	var out2 = strings.TrimSuffix(string(out1), "\n")
	//converts the string into an int
	var num, _ = strconv.Atoi(out2)

	if err != nil {
		errorTput()
		errorExit()
	}

	return num
}

func fillSpaces(num int) string {

	var spaces string

	for i := 1; i <= num; i++ {

		spaces += " "

	}

	return spaces
}

// add split() from go-reloaded

func splitWhiteSpaces(s string) []string {

	var test = ""
	var size = 1
	var index = 0
	var count = 0
	var space rune

	for _, v := range s {

		if (space != ' ' && space != '\t' && space != '\n') || (v != ' ' && v != '\t' && v != '\n') {

			test += string(v)

		}

	}

	for i := 0; i < len(test); i++ {

		if test[i] == ' ' || test[i] == '\t' || test[i] == '\n' {
			size++
		}

	}

	var result = make([]string, size)

	for i := 0; i < len(test); i++ {

		if test[i] == ' ' || test[i] == '\t' || test[i] == '\n' {

			result[count] = result[count] + test[index:i]
			index = i + 1
			count++

		}

	}

	result[count] += test[index:]

	return result

}
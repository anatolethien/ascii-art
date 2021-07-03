package main

import (
	// "github.com/anatolethien/ascii-art/pkg/go-reloaded"
	"fmt"
	"github.com/anatolethien/ascii-art/pkg/ascii-art"
)

func main() {
	ascii.GetFile("standard")
	fmt.Println(ascii.Classic())
	// fmt.Println(reloaded.Atoi("00012345"))
	// for i, v := range ascii.GetArgs() {
	//     fmt.Printf("[%d] = %s\n", i, v)
	// }
	// fmt.Println(ascii.GetArgs())
}

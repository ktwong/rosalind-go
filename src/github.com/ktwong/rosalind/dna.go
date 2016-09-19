package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	inputPath := os.Args[1]
	inputFile, err := ioutil.ReadFile(inputPath)
	check(err)
	input := string(inputFile)

	aa := strings.Count(input, "A")
	cc := strings.Count(input, "C")
	gg := strings.Count(input, "G")
	tt := strings.Count(input, "T")

	fmt.Fprintf(os.Stdout, "%d %d %d %d\n", aa, cc, gg, tt)

}

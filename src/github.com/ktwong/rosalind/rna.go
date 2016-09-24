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

	output := strings.Replace(input, "T", "U", -1)

	fmt.Fprintf(os.Stdout, "%s\n", output)

}

package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

var total = 0
var count = 0
var in string
var stdin = false
var line = false //vs all

func init() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		os.Exit(1)
	}
	in = flag.Arg(0)
	switch flag.Arg(1) {
	case "line", "lines":
		line = true
	case "all":
		line = false
	default:
		os.Exit(2)
	}
}

func main() {
	if strings.EqualFold(in, "stdin") {

	} else {

	}
	os.Exit(0)
}

func fromStdin() {
	if line {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
	} else {
		ioutil.ReadAll(os.Stdin)
	}
}

func eval(i ...int) {
	for _, val := range i {
		total += val
	}
	count += len(i)
}

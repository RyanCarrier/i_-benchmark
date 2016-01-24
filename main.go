package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/RyanCarrier/i_-benchmark/inputs"
)

var argv []string
var c inputs.Cfg

func init() {
	flag.Parse()
	argv = flag.Args()
	c = inputs.NewCfg()
}

func main() {
	if err := c.SetupAndRun(argv); err != nil {
		fmt.Println(err.Error())
		inputs.Usage()
	}
	fmt.Println("Total\tCount\n" + strconv.Itoa(c.Total) + "\t" + strconv.Itoa(c.Count))
}

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//Counts keeps track of all the files that are needed to be created
var Counts []int

var add bool

func init() {
	flag.Parse()
	argv := flag.Args()
	switch len(argv) {
	case 0, 1:
		usage()
	case 2:
		c, err := parseSingle(argv[1])
		if err != nil {
			usage()
		}
		Counts = c
	default:
		c, err := parseMultiple(argv[1:])
		if err != nil {
			usage()
		}
		Counts = c
	}
	setupAdd(argv[0])
}

func main() {
	for _, i := range Counts {
		if add {
			createFile(i)
		} else {
			os.Remove(numToFile(i))
		}
	}
}

func setupAdd(s string) {
	switch strings.ToLower(s) {
	case "add":
		add = true
	case "remove":
		add = false
	default:
		usage()
	}
}

func parseSingle(s string) (ints []int, err error) {
	split := strings.Split(s, ",")
	stringsToInts(split)
	return
}

func parseMultiple(ss []string) (ints []int, err error) {
	return stringsToInts(ss)

}

func stringsToInts(s []string) (ints []int, err error) {
	ints = make([]int, len(s))
	for i, a := range s {
		if ints[i], err = strconv.Atoi(a); err != nil {
			return
		}
	}
	return
}

func usage() {
	fmt.Println("./main add/remove 1 [10 100 ...]")
	fmt.Println("\tor")
	fmt.Println("./main add/remove 1[,10,100,...]")
	os.Exit(1)
}

func numToFile(i int) string {
	return "BenchFile" + strconv.Itoa(i) + ".in"
}

func createFile(ints int) string {
	rand.Seed(1234)
	temp := 0
	filename := numToFile(ints)
	os.Remove(filename)
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	randTop := ints / 2
	if randTop <= 0 {
		randTop = 1
	}
	//finalString := ""
	for i := 0; i < ints; i++ {
		temp++
		append := strconv.Itoa(rand.Intn(100))
		append += " "
		if rand.Intn(randTop) < temp {
			temp = 0
			append += "\n"
		}
		//finalString += append
		file.WriteString(append)
	}
	file.Close()
	return filename
}

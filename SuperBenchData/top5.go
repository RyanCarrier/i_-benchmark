package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Bench keeps track of benchmark results
type Bench struct {
	Name    string
	NsPerOp int
	Input   string
}

//Benchs keeps track of a set of Bench's
type Benchs struct {
	Bs []Bench
}

var in = []string{"Bench300B", "Bench30kB", "Bench3MB", "Bench30MB", "Bench300MB"}
var out = "Results"
var sets = []string{"PASS", "CAT", "FILE"}

//3 parse methods, 6 read methods, 3 input methods
var totalBenchs = 3 * 6 * 3

func init() {
	for _, f := range in {
		_, err := os.Stat(f)
		if err != nil {
			fmt.Println("Error finding input file;" + f)
			os.Exit(1)
		}
	}
}

func main() {
	//os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	//outFile, _ := os.Open(out)
	for _, f := range in {
		out, _ := ioutil.ReadFile(f)
		B := fieldsToStructs(splitByInputs(string(out)))
		Bs := Benchs{Bs: B}
		fmt.Print(f + "\n")
		fmt.Print("Mean\tMedian\n")
		fmt.Print(strconv.Itoa(Bs.mean()) +
			"\t" + strconv.Itoa(Bs.median()) + "\n")
		fmt.Print("Top 5;\n")
		for _, b := range Bs.fast5() {
			fmt.Print(b.toString() + "\n")
		}
		fmt.Print("Slowest 5;\n")
		for _, b := range Bs.slow5() {
			fmt.Print(b.toString() + "\n")
		}
		fmt.Print("\n")
		//fmt.Println(Bs.mean())
	}
}

func (b Bench) toString() string {
	return b.Input + "\t" + b.Name + "\t" + strconv.Itoa(b.NsPerOp) +
		"\tns/op"
}

func (b Benchs) mean() int {
	sum := int64(0)
	for _, bench := range b.Bs {
		sum += int64(bench.NsPerOp)
	}
	return int(sum / int64(len(b.Bs)))
}

func (b Benchs) median() int {
	sort.Sort(b)
	return b.Bs[len(b.Bs)/2].NsPerOp
}

func (b Benchs) fast5() []Bench {
	sort.Sort(b)
	return b.Bs[:5]
}

func (b Benchs) slow5() []Bench {
	sort.Sort(b)
	return b.Bs[len(b.Bs)-5:]
}

//Len length
func (b Benchs) Len() int {
	return len(b.Bs)
}

//Less is it less than
func (b Benchs) Less(i, j int) bool {
	return b.Bs[i].NsPerOp < b.Bs[j].NsPerOp
}

//Swap swaps them
func (b Benchs) Swap(i, j int) {
	b.Bs[i], b.Bs[j] = b.Bs[j], b.Bs[i]
}

func fieldsToStructs(s [][]string) []Bench {
	Result := make([]Bench, totalBenchs)
	i := 0
	for j, set := range s {
		for _, l := range set {
			//fmt.Println(l)
			if strings.TrimSpace(l) == "" {
				continue
			}
			fields := strings.Fields(l)
			if len(fields) < 3 {
				fmt.Println("ERROR:")
				fmt.Println(fields)
			}
			ns, _ := strconv.Atoi(fields[2])
			Result[i] = Bench{
				Name:    fields[0],
				NsPerOp: ns,
				Input:   sets[j],
			}
			i++
		}
	}
	return Result
}
func splitByInputs(s string) [][]string {
	total := make([][]string, 3)
	for i := 0; i < 3; i++ {
		total[i] = make([]string, 3*6)
	}
	index := 0
	j := 0

	for i, l := range strings.Split(s, "\n") {
		if i == 0 {
			continue
		}
		trimmed := strings.TrimSpace(l)
		switch trimmed {
		case "":
			continue
		case sets[0]:
			index = 0
			j = 0
		case sets[1]:
			index = 1
			j = 0
		case sets[2]:
			index = 2
			j = 0
		default:
			//fmt.Println(l)
			total[index][j] = l
			j++
		}
	}
	return total
}

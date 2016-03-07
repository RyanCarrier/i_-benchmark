package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

//RM, read methods
var RM = []string{"ioutil", "ioutilmanual", "bufioscanint", "bufioscanlines", "bufioline", "bufioall"}
var PM = []string{"fmtscan", "scan", "splitstrconv"}

//0; pass 1; cat 2; file
var Input int
var TestSize int
var Filename string

func init() {
	TestSize = getTestSize()
	if TestSize == -1 {
		fmt.Fprint(os.Stderr, "TestSize not set correctly")
	}
	Filename = numToFile(TestSize)
	Input = getInputType()
	if Input == -1 {
		fmt.Fprint(os.Stderr, "Input not set correctly")
		os.Exit(1)
	}
}

//ioutil
func BenchmarkRM0PM0(b *testing.B) { benchmark(b, 0, 0) }
func BenchmarkRM0PM1(b *testing.B) { benchmark(b, 0, 1) }
func BenchmarkRM0PM2(b *testing.B) { benchmark(b, 0, 2) }

//ioutilmanual
func BenchmarkRM1PM0(b *testing.B) { benchmark(b, 1, 0) }
func BenchmarkRM1PM1(b *testing.B) { benchmark(b, 1, 1) }
func BenchmarkRM1PM2(b *testing.B) { benchmark(b, 1, 2) }

//bufioscanint
func BenchmarkRM2PM0(b *testing.B) { benchmark(b, 2, 0) }
func BenchmarkRM2PM1(b *testing.B) { benchmark(b, 2, 1) }
func BenchmarkRM2PM2(b *testing.B) { benchmark(b, 2, 2) }

//bufioscanlines
func BenchmarkRM3PM0(b *testing.B) { benchmark(b, 3, 0) }
func BenchmarkRM3PM1(b *testing.B) { benchmark(b, 3, 1) }
func BenchmarkRM3PM2(b *testing.B) { benchmark(b, 3, 2) }

//bufioline
func BenchmarkRM4PM0(b *testing.B) { benchmark(b, 4, 0) }
func BenchmarkRM4PM1(b *testing.B) { benchmark(b, 4, 1) }
func BenchmarkRM4PM2(b *testing.B) { benchmark(b, 4, 2) }

//bufioall
func BenchmarkRM5PM0(b *testing.B) { benchmark(b, 5, 0) }
func BenchmarkRM5PM1(b *testing.B) { benchmark(b, 5, 1) }
func BenchmarkRM5PM2(b *testing.B) { benchmark(b, 5, 2) }

func benchmark(b *testing.B, rmi, pmi int) {
	var c *exec.Cmd
	switch Input {
	//Pass
	case 0:
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c = exec.Command("./main", "stdin", RM[rmi], PM[pmi], "<", Filename)
			c.Run()
		}
		//cat
	case 1:
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c = exec.Command("cat", Filename, "|", "./main", "stdin", RM[rmi], PM[pmi])
			c.Run()
		}
		//File
	case 2:
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c = exec.Command("./main", Filename, RM[rmi], PM[pmi])
			c.Run()
		}
	default:
		b.Error("Something went wrong")
	}

}

func numToFile(i int) string {
	return "BenchFile" + strconv.Itoa(i) + ".in"
}

func getInputType() int {
	p, err := ioutil.ReadFile("BenchInput.in")
	if err != nil {
		return -1
	}
	switch strings.ToLower(string(p)) {
	case "pass":
		return 0
	case "cat":
		return 1
	case "file":
		return 2
	default:
		return -1
	}
}

func getTestSize() int {
	p, err := ioutil.ReadFile("BenchSize.in")
	if err != nil {
		return -1
	}
	i, err := strconv.Atoi(string(p))
	//fmt.Println(err.Error())
	if err != nil {
		return -1
	}
	return i
}

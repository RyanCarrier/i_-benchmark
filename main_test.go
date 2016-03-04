package main

import (
	"os/exec"
	"strconv"
	"testing"
)

var ReadMethods = []string{"bufio", "ioutil", "ioutilmanual", "bufioscanint", "bufioscanlines", "bufioline", "bufioall"}
var TestFileNumbers = []int{1, 5, 50, 10000}

func BenchmarkIoutil100(b *testing.B) {

}

func BenchmarkIoutilManual100(b *testing.B) {

}

func BenchmarkBufioScanint100(b *testing.B) {

}

func BenchmarkBufioScanlines100(b *testing.B) {

}

func BenchmarkBufioLine100(b *testing.B) {

}

func BenchmarkBufioAll100(b *testing.B) {

}

func benchmark(ints int, b *testing.B, ReadMethod, BufioReadStyle, ParseMethod string) {
	filename := numToFile(ints)
	var c *exec.Cmd
	if ReadMethod == "bufio" {
		for n := 0; n < b.N; n++ {
			c = exec.Command("./main", filename, ReadMethod, BufioReadStyle, ParseMethod)
			b.ResetTimer()
			c.Run()
		}
	} else {
		for n := 0; n < b.N; n++ {
			c = exec.Command("./main", filename, ReadMethod, ParseMethod)
			b.ResetTimer()
			c.Run()
		}
	}
}

func numToFile(i int) string {
	return "BenchFile" + strconv.Itoa(i) + ".in"
}

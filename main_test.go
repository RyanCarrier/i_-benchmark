package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func benchmark(ints int, b *testing.B, ReadMethod, BufioReadStyle, ParseMethod string) {

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

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var total = 0
var count = 0
var inFile string     // stdin/file
var readMethod string // bufio/ioutil/readfile
var readStyle string  // line/all
var argv []string

func init() {
	flag.Parse()
	if len(flag.Args()) != 3 {
		os.Exit(1)
	}
	argv = flag.Args()
}

func init() {
	inFile = argv[0]
	readMethod = argv[1]
	readStyle = argv[2]
}

func main() {
	var f *os.File
	if argv[0] == "stdin" {
		f = os.Stdin
	} else {
		f, _ = os.Open(argv[0])
		defer f.Close()
	}
	fromFile(f)
}

func fromFile(f *os.File) {
	switch readMethod {
	case "bufio":
		fromBufio(f)
	case "ioutil":
		fromIoutil(f)
	case "readFile":
		fromReadFile(f)
	default:
		os.Exit(3) //SOMETHING WRONG
	}
}

//ReadFile does the same as
func fromReadFile(f *os.File) {
	var n int64
	if fi, err := f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	ioutilReadFile(f, n)
}

func fromIoutil(f *os.File) {
	ioutil.ReadAll(f) //reallocate buffer

}

func fromBufio(f *os.File) {
	var p []byte
	reader := bufio.NewReader(f)
	switch readStyle {
	case "line":
		p, _ = reader.ReadBytes('\n') //readstring calls readbytes
	case "all":
		s, _ := f.Stat()
		p = make([]byte, s.Size()+bytes.MinRead)
		reader.Read(p)
	default:
		os.Exit(4) //SOMETHING WRONG
	}
}

func usage() {
	fmt.Println("main stdin/infile bufio/ioutil all/line ")
}
func convertLine(s []byte) {

}

//FROM READALL IOUTIL
// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
//rcarrier: We can't use the usual ioutil.ReadFile as we can't pass in stdin
// that way. Also we don't want to have to re-open the input file.
func ioutilReadFile(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func eval(i ...int) {
	for _, val := range i {
		total += val
	}
	count += len(i)
}

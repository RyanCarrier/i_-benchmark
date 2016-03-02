package methodTest

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//ReadFromIoutilReadAll does the same as ioutil.ReadFile without re-opening the file.
func ReadFromIoutilReadAll(f *os.File) ([]int, error) {
	var n int64
	if fi, err := f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	b, err := ioutilReadAll(f, n)
	if err != nil {
		return []int{}, errors.New("Error reading by ioutilReadAll, ReadFromIoutilReadAll; " + err.Error())
	}
	return StringToInts(string(b)), nil
}

//FromIoutil reads a file using ioutil(ReadAll)
func FromIoutil(f *os.File) ([]int, error) {
	//MAYBE CLOSE THEN REOPEN
	r, err := ioutil.ReadAll(f) //reallocate buffer
	if err != nil {
		return []int{}, errors.New("Error reading by ioutil.ReadAll, FromIoutil; " + err.Error())
	}
	return StringToInts(string(r)), nil
}

//FromBufioReadBytes reads
func FromBufioReadBytes(f *os.File) (result []int, err error) {
	var p []byte
	reader := bufio.NewReader(f)
	//readstring calls readbytes
	for p, err = reader.ReadBytes('\n'); err == nil; p, err = reader.ReadBytes('\n') {
		result = append(result, StringToInts(string(p))...)
		//if err = c.EvaluateLine(p); err != nil {
		//	return []int{}, errors.New("Error reading by line, FromBufio; " + err.Error())
		//}
	}
	if err != nil && err != io.EOF {
		return
	}
	err = nil
	result = append(result, StringToInts(string(p))...)
	return
}

//FromBufioFmtScan read
func FromBufioFmtScan(f *os.File) (result []int, err error) {
	var i int
	reader := bufio.NewReader(f)
	for _, err := fmt.Fscan(reader, &i); err == nil; _, err = fmt.Fscan(reader, &i) {
		result = append(result, i)
	}
	if err != nil && err != io.EOF {
		err = errors.New("Error reading by scanint, FromBufio; " + err.Error()) //NOTE this will be something, even if it runs correctly...
		return
	}
	err = nil
	return
}

//FromBufioScanner read
func FromBufioScanner(f *os.File) (result []int, err error) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, StringToInts(string(scanner.Bytes()))...)
	}
	return
}

//FromBufioReaderRead read
func FromBufioReaderRead(f *os.File) (result []int, err error) {
	reader := bufio.NewReader(f)
	s, err := f.Stat()
	if err != nil {
		err = errors.New("Error reading by all, FromBufio; " + err.Error())
		return []int{}, err
	}
	p := make([]byte, s.Size()) //+bytes.MinRead)
	reader.Read(p)
	fmt.Println(string(p))
	result = StringToInts(string(p))
	fmt.Println(result)
	return result, nil
}

//FROM READALL IOUTIL
// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
//rcarrier: We can't use the usual ioutil.ReadFile as we can't pass in stdin
// that way. Also we don't want to have to re-open the input file.
func ioutilReadAll(r io.Reader, capacity int64) (b []byte, err error) {
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

//StringToInts converts a string to an array of ints
func StringToInts(s string) []int {
	fields := strings.Fields(s)
	result := make([]int, len(fields))
	for index, f := range fields {
		i, _ := strconv.Atoi(f)
		result[index] = i
	}
	return result
}

//Calc returns amount of numbers and their total
func Calc(values []int) (count, total int) {
	for _, i := range values {
		total += i
	}
	return len(values), total
}

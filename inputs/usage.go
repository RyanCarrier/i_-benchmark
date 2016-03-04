package inputs

import (
	"fmt"
	"os"
)

//Usage prints the help/usage options
func Usage() {
	fmt.Println("Usage:")
	printt("main SourceFile ReadMethod ParseMethod")
	usageSourceFile()
	usageReadMethod()
	os.Exit(1)
}
func usageSourceFile() {
	fmt.Println("SourceFile;")
	printt("stdin\ttakes input from stdin")
	printt("file\t(any input file)")
}

func usageReadMethod() {
	fmt.Println("ReadMethod;")
	printt("ioutil\tioutil is readfile with extra file open")
	printt("ioutilmanual\tuses ioutils internal readfile, without re-opening the file")
	printt("bufioscanint\tuses bufio reader and fmt.Fscan to scan individual ints")
	printt("bufioscanlines\tuses default bufio scanner to parse line by line")
	printt("bufioline\tuses bufio reader to read line by line")
	printt("bufioall\tuses bufio reader to create a perfect sized buffer and read all content")
}

func printt(s string) {
	fmt.Println("\t" + s)
}

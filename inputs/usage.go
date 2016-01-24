package inputs

import (
	"fmt"
	"os"
)

//Usage prints the help/usage options
func Usage() {
	fmt.Println("Usage:")
	printt("main SourceFile ReadMethod BufioReadStyle ParseMethod")
	fmt.Println(`In the case of "ReadMethod : bufio", "BufioReadStyle : scanint", ParseMethod is not needed.`)
	usageSourceFile()
	usageReadMethod()
	usageBufioReadStyle()
	os.Exit(1)
}
func usageSourceFile() {
	fmt.Println("SourceFile;")
	printt("stdin")
	printt("file\t(any input file)")
}

func usageReadMethod() {
	fmt.Println("ReadMethod;")
	printt("bufio\tMust specify a BufioReadStyle")
	printt("ioutil\tioutil is readfile with extra file open")
	printt("readfile\tuses ioutils internal readfile")
}

func usageBufioReadStyle() {
	fmt.Println("BufioReadStyle;")
	printt("scanint\tscans by each int found, ParseMethod not used")
	printt("scanlines\t scans by each line, then uses ParseMethod to evaluate line")
}
func printt(s string) {
	fmt.Println("\t" + s)
}

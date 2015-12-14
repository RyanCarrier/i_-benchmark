package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	argv := flag.Args()
	f, _ := os.Open(argv[0])
	fs, _ := f.Stat()
	fmt.Println("Filesize;")
	fmt.Println(fs.Size())
	f.Close()
	
	fs, _ = os.Stdin.Stat()
	fmt.Println("Stdinsize;")
	fmt.Println(fs.Size())

}

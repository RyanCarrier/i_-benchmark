package inputs

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

//Cfg tells the package what configurations are being used.
type Cfg struct {
	//Total keeps track of the total amount when the numbers are summed.
	Total int
	//Count keeps track of the amount of numbers processed.
	Count int
	//SourceFile is the input file, a file or stdin
	SourceFile string // stdin/file
	//ReadMethod determines which method is used to retrieve the data from file
	ReadMethod string // bufio/ioutil/readfile
	//BufioReadStyle is what type of reading should be used when bufio read is selected
	BufioReadStyle string // scanint/scanlines/line/all
	//ParseMethod decides which way the integers are passed from byte slices
	ParseMethod string // scan/fmtscan/splitstrconv
	//f is the pointer to the file to read from, file or stdin
	f *os.File
}

func usage() {
	fmt.Println("Usage:")
	printt("main SourceFile ReadMethod BufioReadStyle ParseMethod")
	fmt.Println(`In the case of "ReadMethod : bufio", "BufioReadStyle : scanint", ParseMethod is not needed.`)
	os.Exit(29)
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

//NewCfg gets a new config struct.
func NewCfg() *Cfg {
	return &Cfg{
		Total: 0,
		Count: 0,
	}
}

//Exec is the main function of this package,
// runs through the specified input file with parameters set in Cfg.
func (c *Cfg) Exec() error {
	switch c.SourceFile {
	case "", "stdin":
		c.f = os.Stdin
	default:
		c.f, _ = os.Open(c.SourceFile)
		defer c.f.Close()
	}
	err := c.ReadFromFile()
	//defer will run before returned? TODO: Check this
	return err
}

//ReadFromFile forwards the data reading job off to the correct function.
//Alos note stdin is considered a file in this case
func (c *Cfg) ReadFromFile() error {
	switch c.ReadMethod {
	case "bufio":
		return c.fromBufio()
	case "ioutil":
		return c.fromIoutil()
	case "readFile":
		return c.ReadFromIoutilReadAll()
	default:
		return errors.New("ReadMethod set incorrectly; " + c.ReadMethod)
	}
}

//ReadFromIoutilReadAll does the same as ioutil.ReadFile without re-opening the file.
func (c *Cfg) ReadFromIoutilReadAll() error {
	var n int64
	if fi, err := c.f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	b, err := ioutilReadAll(c.f, n)
	if err != nil {
		return err
	}
	return c.EvaluateAll(b)
}

func (c *Cfg) fromIoutil() error {
	//MAYBE CLOSE THEN REOPEN
	r, err := ioutil.ReadAll(c.f) //reallocate buffer
	if err != nil {
		return err
	}
	return c.EvaluateAll(r)
}

func (c *Cfg) fromBufio() error {
	var p []byte
	var err error
	switch c.BufioReadStyle {
	case "line":
		reader := bufio.NewReader(c.f)
		//readstring calls readbytes
		for p, err = reader.ReadBytes('\n'); err == nil; p, err = reader.ReadBytes('\n') {
			if err = c.EvaluateLine(p); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		return c.EvaluateLine(p)
	case "scanint":
		var i int
		reader := bufio.NewReader(c.f)
		for _, err := fmt.Scan(reader, &i); err == nil; _, err = fmt.Scan(reader, &i) {
			c.eval(i)
		}
		return err //NOTE this will be something, even if it runs correctly...
	case "scanlines":
		scanner := bufio.NewScanner(c.f)
		for scanner.Scan() {
			err = c.EvaluateLine(scanner.Bytes())
			return err
		}
	case "all":
		reader := bufio.NewReader(c.f)
		s, err := c.f.Stat()
		if err != nil {
			return err
		}
		p = make([]byte, s.Size()+bytes.MinRead)
		reader.Read(p)
		return c.EvaluateAll(p)
	default:
		return errors.New("BufioReadStyle incorrectly set; " + c.BufioReadStyle)
	}
	return nil
}

//EvaluateAll evaluates everything passed in, multiple lines.
func (c *Cfg) EvaluateAll(s []byte) error {
	var i int
	switch c.ParseMethod {
	case "fmtscan":
		reader := bytes.NewReader(s)
		for _, err := fmt.Scan(reader, &i); err == nil; _, err = fmt.Scan(reader, &i) {
			if err != nil {
				return err
			}
			c.eval(i)
		}
	case "scan":
		scanner := bufio.NewScanner(bytes.NewReader(s))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text()) // string(scanner.Bytes())
			if err != nil {
				return err
			}
			c.eval(i)
		}
	case "splitstrconv":
		fields := strings.Fields(string(s))
		for _, f := range fields {
			i, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			c.eval(i)
		}
	default:
		return errors.New("ParseMethod not set correctly; " + c.ParseMethod)
	}
	return nil
}

//EvaluateLine converts the line from a byte slice to integers, then adding
// them to Count and sum
func (c *Cfg) EvaluateLine(s []byte) error {
	var i int
	switch c.ParseMethod {
	case "fmtscan":
		reader := bytes.NewReader(s)
		for _, err := fmt.Scan(reader, &i); err == nil; _, err = fmt.Scan(reader, &i) {
			if err != nil {
				return err
			}
			c.eval(i)
		}
	case "scan":
		scanner := bufio.NewScanner(bytes.NewReader(s))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text()) // string(scanner.Bytes())
			if err != nil {
				return err
			}
			c.eval(i)
		}
	case "splitstrconv":
		fields := strings.Fields(string(s))
		for _, f := range fields {
			i, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			c.eval(i)
		}
	default:
		return errors.New("ParseMethod not set correctly; " + c.ParseMethod)
	}
	return nil
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

func (c *Cfg) eval(i ...int) {
	for _, val := range i {
		c.Total += val
	}
	c.Count += len(i)
}

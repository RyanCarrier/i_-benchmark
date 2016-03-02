package inputs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

//SetupAndRun sets the cfg struct correctly, printing usage if errors.
func (c *Cfg) SetupAndRun(argv []string) error {
	if len(argv) < 3 || len(argv) > 4 {
		return errors.New("Too little or too many arguments.")
	}
	c.SourceFile = strings.ToLower(argv[0])
	c.ReadMethod = strings.ToLower(argv[1])
	if len(argv) == 4 {
		if c.ReadMethod != "bufio" {
			return errors.New("To set BufioReadStyle, ReadMethod must be bufio.")
		}
		c.BufioReadStyle = strings.ToLower(argv[2])
		c.ParseMethod = strings.ToLower(argv[3])
	} else {
		c.ParseMethod = strings.ToLower(argv[2])
	}
	return c.Exec()
}

//NewCfg gets a new config struct.
func NewCfg() Cfg {
	return Cfg{
		Total: 0,
		Count: 0,
	}
}

//Exec is the main function of this package,
// runs through the specified input file with parameters set in Cfg.
func (c *Cfg) Exec() error {
	var err error
	c.Total = 0
	c.Count = 0
	switch c.SourceFile {
	case "", "stdin":
		c.f = os.Stdin
	default:
		c.f, err = os.Open(c.SourceFile)
		if err != nil {
			return errors.New("Error opening source file; " + err.Error())
		}
		defer c.f.Close()
	}
	err = c.read()
	return err
}

//read forwards the data reading job off to the correct function.
//Alos note stdin is considered a file in this case
//read is private as files NEED to be opened before running through
func (c *Cfg) read() error {
	switch strings.ToLower(c.ReadMethod) {
	case "bufio":
		return c.ReadBufio()
	case "ioutilmanual":
		return c.ReadIoutilReadAllManual()
	case "ioutil":
		return c.ReadIoutilReadAll()
	default:
		return errors.New("ReadMethod set incorrectly; " + c.ReadMethod)
	}
}

//ReadIoutilReadAllManual does the same as ioutil.ReadFile without re-opening the file.
func (c *Cfg) ReadIoutilReadAllManual() error {
	var n int64
	if fi, err := c.f.Stat(); err == nil {
		// Don't preallocate a huge buffer, just in case.
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	}
	b, err := ioutilReadAll(c.f, n)
	if err != nil {
		return errors.New("Error reading by ioutilReadAll, ReadFromIoutilReadAll; " + err.Error())
	}
	return c.EvaluateAll(b)
}

//ReadIoutilReadAll reads a file using ioutil(ReadAll)
func (c *Cfg) ReadIoutilReadAll() error {
	//MAYBE CLOSE THEN REOPEN
	r, err := ioutil.ReadAll(c.f) //reallocate buffer
	if err != nil {
		return errors.New("Error reading by ioutil.ReadAll, ReadIoutilReadAll; " + err.Error())
	}
	return c.EvaluateAll(r)
}

//ReadBufio evaluates which bufio method to use when reading the file, then
//does it.
func (c *Cfg) ReadBufio() error {
	var p []byte
	var err error
	switch c.BufioReadStyle {
	case "line":
		reader := bufio.NewReader(c.f)
		//readstring calls readbytes
		for p, err = reader.ReadBytes('\n'); err == nil; p, err = reader.ReadBytes('\n') {
			if err = c.EvaluateLine(p); err != nil {
				return errors.New("Error reading by line, ReadBufio; " + err.Error())
			}
		}
		if err != nil && err != io.EOF {
			return err
		}
		return c.EvaluateLine(p)
	case "scanint":
		var i int
		reader := bufio.NewReader(c.f)
		for _, err := fmt.Fscan(reader, &i); err == nil; _, err = fmt.Fscan(reader, &i) {
			c.eval(i)
		}
		if err != nil && err != io.EOF {
			return errors.New("Error reading by scanint, ReadBufio; " + err.Error()) //NOTE this will be something, even if it runs correctly...
		}
		return nil
	case "scanlines":
		scanner := bufio.NewScanner(c.f)
		for scanner.Scan() {
			err = c.EvaluateLine(scanner.Bytes())
			if err != nil {
				return errors.New("Error reading by scanlines, ReadBufio; " + err.Error())
			}
		}
		return nil
	case "all":
		reader := bufio.NewReader(c.f)
		s, err := c.f.Stat()
		if err != nil {
			return errors.New("Error reading by all, ReadBufio; " + err.Error())
		}
		p = make([]byte, s.Size()) //+bytes.MinRead)
		reader.Read(p)
		//fmt.Println(string(p))
		return c.EvaluateAll(p)
	default:
		return errors.New("BufioReadStyle incorrectly set; " + c.BufioReadStyle)
	}
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

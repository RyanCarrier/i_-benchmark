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
	total       int
	count       int
	inFile      string // stdin/file
	readMethod  string // bufio/ioutil/readfile
	readStyle   string // scan/scanlines/line/all
	parseMethod string // scan/fmtscan/splitstrconv
	f           *os.File
}

//NewCfg gets a new config struct.
func NewCfg() *Cfg {
	return &Cfg{
		total: 0,
		count: 0,
	}
}

//Exec is the main function of this package,
// runs through the specified input file with parameters set in Cfg.
func (c *Cfg) Exec() error {
	switch c.inFile {
	case "", "stdin":
		c.f = os.Stdin
	default:
		c.f, _ = os.Open(c.inFile)
		defer c.f.Close()
	}
	return c.fromFile()
}

func (c *Cfg) fromFile() error {
	switch c.readMethod {
	case "bufio":
		return c.fromBufio()
	case "ioutil":
		return c.fromIoutil()
	case "readFile":
		return c.fromReadFile()
	default:
		return errors.New("readMethod set incorrectly; " + c.readMethod)
	}
}

//ReadFile does the same as ioutil.ReadFile without re-opening the file.
func (c *Cfg) fromReadFile() error {
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
	return c.convertAll(b)
}

func (c *Cfg) fromIoutil() error {
	r, err := ioutil.ReadAll(c.f) //reallocate buffer
	if err != nil {
		return err
	}
	return c.convertAll(r)
}

func (c *Cfg) fromBufio() error {
	var p []byte
	var err error
	switch c.readStyle {
	case "line":
		reader := bufio.NewReader(c.f)
		p, err = reader.ReadBytes('\n') //readstring calls readbytes
		if err != nil {
			return err
		}
		return c.convertLine(p)
	case "scan":
		var i int
		reader := bufio.NewReader(c.f)
		for _, err := fmt.Scan(reader, &i); err == nil; _, err = fmt.Scan(reader, &i) {
			c.eval(i)
		}
		return err //NOTE this will be something, even if it runs correctly...
	case "scanlines":
		scanner := bufio.NewScanner(c.f)
		for scanner.Scan() {
			err = c.convertLine(scanner.Bytes())
			if err != nil {
				return err
			}
		}
	case "all":
		reader := bufio.NewReader(c.f)
		s, err := c.f.Stat()
		if err != nil {
			return err
		}
		p = make([]byte, s.Size()+bytes.MinRead)
		reader.Read(p)
		return c.convertAll(p)
	default:
		return errors.New("readStyle incorrectly set; " + c.readStyle)
	}
	return nil
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tmain InputSource ParseMethod ParseByLine")
	fmt.Println("\tmain stdin/infile bufio/ioutil all/line ")
	os.Exit(29)
}
func (c *Cfg) convertAll(s []byte) error {
	var i int
	switch c.parseMethod {
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
		return errors.New("parseMethod not set correctly; " + c.parseMethod)
	}
	return nil
}

func (c *Cfg) convertLine(s []byte) error {
	var i int
	switch c.parseMethod {
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
		return errors.New("parseMethod not set correctly; " + c.parseMethod)
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
		c.total += val
	}
	c.count += len(i)
}

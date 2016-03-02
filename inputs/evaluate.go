package inputs

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//EvaluateAll evaluates everything passed in, multiple lines.
func (c *Cfg) EvaluateAll(s []byte) error {
	var i int
	switch strings.ToLower(c.ParseMethod) {
	case "fmtscan":
		reader := bytes.NewReader(s)
		for _, err := fmt.Fscan(reader, &i); err == nil; _, err = fmt.Fscan(reader, &i) {
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
		if c.ParseMethod == "" {
			return errors.New("ParseMethod not set")
		}
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
		_, err := fmt.Fscan(reader, &i)
		//	fmt.Println(n)
		for err == nil {
			if err != nil {
				return err
			}
			c.eval(i)
			_, err = fmt.Fscan(reader, &i)
			//	fmt.Println(n)
		}
		break
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
		break
	case "splitstrconv":
		fields := strings.Fields(string(s))
		for _, f := range fields {
			i, err := strconv.Atoi(f)
			if err != nil {
				return err
			}
			c.eval(i)
		}
		break
	default:
		if c.ParseMethod == "" {
			return errors.New("ParseMethod not set")
		}
		return errors.New("ParseMethod not set correctly; " + c.ParseMethod)
	}
	return nil
}

func (c *Cfg) eval(i ...int) {
	for _, val := range i {
		c.Total += val
	}
	c.Count += len(i)
}

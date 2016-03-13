# inputs
--
    import "github.com/RyanCarrier/i_-benchmark/inputs"


## Usage

#### func  Usage

```go
func Usage()
```
Usage prints the help/usage options

#### type Cfg

```go
type Cfg struct {
	//Total keeps track of the total amount when the numbers are summed.
	Total int
	//Count keeps track of the amount of numbers processed.
	Count int
	//SourceFile is the input file, a file or stdin
	SourceFile string // stdin/file
	//ReadMethod determines which method is used to retrieve the data from file
	ReadMethod string // bufio/ioutil/readfile/bufioscanint/bufioscanlines/bufioline/bufioall
	//ParseMethod decides which way the integers are passed from byte slices
	ParseMethod string // scan/fmtscan/splitstrconv
}
```

Cfg tells the package what configurations are being used.

#### func  NewCfg

```go
func NewCfg() Cfg
```
NewCfg gets a new config struct.

#### func (*Cfg) EvaluateAll

```go
func (c *Cfg) EvaluateAll(s []byte) error
```
EvaluateAll evaluates everything passed in, multiple lines.

#### func (*Cfg) EvaluateLine

```go
func (c *Cfg) EvaluateLine(s []byte) error
```
EvaluateLine converts the line from a byte slice to integers, then adding them
to Count and sum

#### func (*Cfg) Exec

```go
func (c *Cfg) Exec() error
```
Exec is the main function of this package, runs through the specified input file
with parameters set in Cfg.

#### func (*Cfg) ReadBufioAll

```go
func (c *Cfg) ReadBufioAll() error
```
ReadBufioAll uses a bufio reader to read the whole file

#### func (*Cfg) ReadBufioLine

```go
func (c *Cfg) ReadBufioLine() error
```
ReadBufioLine uses a bufio reader to read line by line

#### func (*Cfg) ReadBufioScanInt

```go
func (c *Cfg) ReadBufioScanInt() error
```
ReadBufioScanInt uses a bufio reader and fmt.Fscan to read int by int

#### func (*Cfg) ReadBufioScanLines

```go
func (c *Cfg) ReadBufioScanLines() error
```
ReadBufioScanLines uses a bufio scanner, scanning the file line by line

#### func (*Cfg) ReadIoutilReadAll

```go
func (c *Cfg) ReadIoutilReadAll() error
```
ReadIoutilReadAll reads a file using ioutil(ReadAll)

#### func (*Cfg) ReadIoutilReadAllManual

```go
func (c *Cfg) ReadIoutilReadAllManual() error
```
ReadIoutilReadAllManual does the same as ioutil.ReadFile without re-opening the
file.

#### func (*Cfg) SetupAndRun

```go
func (c *Cfg) SetupAndRun(argv []string) error
```
SetupAndRun sets the cfg struct correctly, printing usage if errors.

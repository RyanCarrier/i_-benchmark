#i\_benchmark
This repository is used for testing the various ways of inputting and parsing data into golang applications.

##Benchmarking
###Options
There are various options to choose from when benchmarking, amount of inputs (filesize), input type (how the file is brought into the application), test duration (or golang default) and a filter, to only test specific read or parse methods.
####Inputs
100;	297B

1000;	2.9kB

10 000;	29kB

...

10 000 000;		29MB

100 000 000;	297MB
####Input type
- pass

Will direct the file in, for example; `./executable < input.in`

- cat

Will cat the file, the pipe it in, for example; `cat input.in | ./executable`

- file

Will let golang open the file itself using [os.OpenFile](https://golang.org/pkg/os/#OpenFile)

- internal

Internal will use a more traditional benchmarking, not compiling and executing the application as the previous input types do. The results of internal are (and should not be) comparable to the other 3 (pass, cat and file).

####Duration
This setting is the approximate time that each benchmark will be ran for. The default is 1s, and 0 will default to that, SuperBench.sh uses 10s.

####Filter
Leaving the filter blank will run with no filter (all), filtering can be done on the read method and(/or) the parse method.
The format of the filter is RM0PM0, where either RM0 or PM0 can be left out (filtering just Parse or Read methods), with each 0 being 0-5 for RM and 0-2 for PM.

#####Read method
The read method regards how the file goes from the file system, to golangs context, in memory.
- 0 : ioutil
ioutil will use golangs [ioutil.ReadAll](https://golang.org/pkg/io/ioutil#ReadAll) to read all the information from the file. As the file is already opened by default when running the executable, ReadAll will re-open the file. To combat this ioutilmanual manually sets the correct variables and runs as though the file was never opened to begin with;

- 1 : ioutilmanual
ioutilmanual will use golangs [ioutil.ReadAll](https://golang.org/pkg/io/ioutil#ReadAll), but will call it manually. By using everything within the function, except for the opening of the file (as the file is already opened by default in the benchmarking executable).

- 2 : bufioscanint
bufioscanint uses a [bufio.Scanner](https://golang.org/pkg/bufio#Scanner), with [fmt.Fscan](https://golang.org/pkg/fmt#Fscan) to scan for integers. Obviously this method will not work as well in practice as it disregards new lines which are generally important.

- 3 : bufioscanlines
bufioscanlines uses a [bufio.Scanner](https://golang.org/pkg/bufio#Scanner) to scan each line and send it for parsing.

- 4 : bufioline
bufioline uses a [bufio.Reader](https://golang.org/pkg/bufio#Reader) to read up to each new line and send it for parsing each time.

- 5 : bufioall
bufioall uses a [bufio.Reader](https://golang.org/pkg/bufio#Reader) to assign a buffer of the correct size and dump the entire file into it for parsing.

#####Parse method
The parse method is how the information from the file goes from a single piece of data (usually a line) to integers and being parsed. This is irrelevant for some read methods such as bufioscanint, as the data gets read as an integer directly from the file.
- 0 : fmtscan
fmtscan, similar to the read method bufioscanints, uses [fmt.Fscan](https://golang.org/pkg/fmt#Fscan) to scan the string for each integer found.

- 1 : scan
scan uses a [bufio.Scanner](https://golang.org/pkg/bufio#Scanner) to scan for each 'word' (integer) and uses [strconv.Atoi](https://golang.org/pkg/strconv#Atoi) to convert the 'word to an integer.

- 2 : splitstrconv
splitstrconv uses [strings.Fields](https://golang.org/pkg/strings#Fields) to split the string into 'words' or 'fields' and then parses each one individually into [strconv.Atoi](https://golang.org/pkg/strconv#Atoi) to convert into an integer.







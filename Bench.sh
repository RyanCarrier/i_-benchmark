#!/bin/bash
rmpmusage(){
  echo -e "\tRead method (RM) options are;"
  echo "0:ioutil"
  echo "1:ioutilmanual"
  echo "2:bufioscanint"
  echo "3:bufioscanlines"
  echo "4:bufioline"
  echo "5:bufioall"
  echo -e "\tParse method (PM) options are;"
  echo "0:fmtscan"
  echo "1:scan"
  echo "2:splitstrconv"
}
usage(){
  echo -e "./gobench inputSize inputType testDuration [filter]\n"
  echo -e "inputSize can be;"
  echo -e "\tAny integer (1,2,500,1000000...)"
  echo -e "\tThis will specify how many integers to use in the benchmark input file"
  echo -e "inputType can be;"
  echo -e "\tpass\twill pass the information in via \"<\" ex: ./main < inputFile"
  echo -e "\tcat\twill pass the information by cat\'ing then piping"
  echo -e "\tfile\twill make the program open the file itself"
  echo -e "\tinternal\twill test without any compilation, golang will open the file itself."
  echo -e "testDuration is how long each test should last (Seconds);"
  echo -e "\t0; for golang default"
  echo -e "\tAny integer, there are many tests, careful with high numbers."
  echo -e "filter is what filter to use when running the benchmarks;"
  echo -e "\tRM for read method, PM for parse method, with the corresponding integer after"
  echo -e "\tRM options are;"
rmpmusage
  echo "EXAMPLES;"
  echo -e "\tTo benchmark by reading with ioutilmanual and parsing with fmtscan method;"
  echo -e "\t./gobench 100 file RM1PM0"
  echo -e "\tTo benchmark by reading with all methods and parsing with splitstrconv method;"
  echo -e "\t./gobench 1234 cat PM2"

  exit 1
}


set -e
#INPUT VALIDATION
if [ $# -lt 3 -o $# -gt 4 ];then
  usage
fi

if [[ ! $1 =~ ^-?[0-9]+$ ]];then
		echo "inputSize non-integer;"
		usage
		exit 1
fi

case $2 in 
		"pass"|"cat"|"file"|"internal")
				;;
		*) 
				echo "INPUTTYPE ERROR;"
				usage
				;;
esac

if [[ ! $3 =~ ^-?[0-9]+$ ]];then
		echo "testDuration non-integer;"
		usage
		exit 1
fi

R1="^PM[0-9]$"
R2="^RM[0-9]$"
R3="^RM[0-9]PM[0-9]$"
FILTER=""
if [ $# -eq 4 ]; then
  if [[ $4 =~ $R1 || $4 =~ $R2 || $4 =~ $R3 ]];then
    FILTER=$4
  else
    echo "Test filter does not match required format."
    echo $4
    echo "Filter must be be an integer variation of;"
    echo -e "\tPM0\n\tRM0\n\tPM0RM0"
    exit 2
  fi
fi
#END INPUT VALIDATION

INPUTFILE="BenchInput.in"
SIZEFILE="BenchSize.in"
echo "Building setup..."
go build benchSetup/main.go
mv ./main ./setup
go build main.go
echo "Building test files..."
./setup add $1
echo -n $1 > $SIZEFILE
echo -n $2 > $INPUTFILE
cp BenchFile* $SIZEFILE ./inputs/

rmpmusage
echo "Benching..."
BENCHTIME=""
if [ $3 -gt 0 ];then
	BENCHTIME="-benchtime=$3s"
fi
if [ $2 = "internal" ];then
		go test ./inputs -bench=".*$FILTER.*" $BENCHTIME
	else
		go test -bench=".*$FILTER.*" $BENCHTIME
fi
echo ""
echo "Cleaning..."
#./setup remove $1
rm BenchFile* inputs/BenchFile* inputs/$SIZEFILE $SIZEFILE $INPUTFILE ./setup ./main

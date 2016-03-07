#!/bin/bash
FOLDER="SuperBench"
FILENAME="Bench"
B1="300B"
B2="30kB"
B3="3MB"
B4="30MB"
B5="300MB"
NAMES=("300B" "30kB" "3MB" "30MB" "300MB")
INTS=(100 10000 100000 10000000 100000000)

echo "TEST $B1" 
LEN=${#NAMES}
for i in `seq 0 $LEN`;do	
		FILE=$FOLDER/$FILENAME${NAMES[i]}

	echo "TESTING ${NAMES[i]}"
	echo "TEST ${NAMES[i]}">$FILE
	echo "PASS">>$FILE
	./gobench ${INTS[i]} pass 10 2>/dev/null |grep Benchmark >> $FILE
	echo "CAT">>$FILE
	./gobench ${INTS[i]} cat 10 2>/dev/null |grep Benchmark >> $FILE
	echo "FILE">>$FILE
	./gobench ${INTS[i]} file 10 2>/dev/null |grep Benchmark >> $FILE
done



echo "Results in ./SuperBench/"

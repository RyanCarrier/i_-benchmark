#!/bin/bash
set -e
echo -e "(very) approximate running time is 1 hour.\nIf canceling early please run ./ManualClean.sh to clean up the benchmarking files.\n"
FOLDER="SuperBench"
FILENAME="Bench"
B1="300B"
B2="30kB"
B3="3MB"
B4="30MB"
B5="300MB"
NAMES=("300B" "30kB" "3MB" "30MB" "300MB")
INTS=(100 10000 100000 10000000 100000000)
mkdir $FOLDER

echo "TEST $B1" 
LEN=${#NAMES}
for i in `seq 0 $LEN`;do	
	FILE="$FOLDER/$FILENAME${NAMES[i]}"

	echo "TESTING ${NAMES[i]}"
	echo "TEST ${NAMES[i]}">$FILE
	echo "PASS">>$FILE
	./Bench.sh ${INTS[i]} pass 10 2>/dev/null |grep Benchmark >> $FILE
	echo "CAT">>$FILE
	./Bench.sh ${INTS[i]} cat 10 2>/dev/null |grep Benchmark >> $FILE
	echo "FILE">>$FILE
	./Bench.sh ${INTS[i]} file 10 2>/dev/null |grep Benchmark >> $FILE
done
echo -e "#!/bin/bash\ngo run ../../SuperBenchData/top5.go" >> "$FOLDER/Top5.sh"


echo "Results in ./$FOLDER"
echo "Run Top5.sh from within ./$FOLDER to see useful statistics and the 5 fastest and slowest combinations for each file size."

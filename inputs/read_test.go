package inputs

import (
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
)

var ReadMethods = []string{"bufio", "ioutil", "ioutilmanual", "bufioscanint", "bufioscanlines", "bufioline", "bufioall"}
var TestFileNumbers = []int{1, 5, 50, 10000}
var Totals []int64

type Test struct {
	in    string
	count int
	total int
}

func TestReadManual(t *testing.T) {
	Tests := []Test{
		Test{
			`1 2 0`,
			3,
			3,
		},
		Test{
			`1 2 2
1 1 1
1000
10000`,
			8,
			11008,
		},
		Test{
			`1 1 1
		1 1 1
		1000
		1`,
			8,
			1007,
		},
	}
	for i, test := range Tests {
		for _, rm := range ReadMethods {
			testManual(t, rm, i, test)
		}
	}
}

func testManual(t *testing.T, ReadMethod string, i int, test Test) {
	ioutil.WriteFile("testdata.in", []byte(test.in), 0666)
	defer os.Remove("testdata.in")
	cfg := NewCfg()
	cfg.ParseMethod = "splitstrconv"
	cfg.SourceFile = "testdata.in"
	cfg.ReadMethod = ReadMethod
	err := cfg.Exec()
	check2(i, t, int64(cfg.Count), int64(cfg.Total), err, int64(test.count), int64(test.total), nil, ReadMethod)

}

func TestRead(t *testing.T) {
	setupTestFiles()
	defer removeTestFiles()
	for _, rm := range ReadMethods {
		test(t, rm, "")
	}
}

func test(t *testing.T, ReadMethod, BufioReadMethod string) {
	for i, Count := range TestFileNumbers {
		cfg := NewCfg()
		cfg.ParseMethod = "splitstrconv"
		cfg.SourceFile = numToFile(Count)
		cfg.ReadMethod = ReadMethod
		err := cfg.Exec()
		check2(i, t, cfg.Count, int64(cfg.Total), err, Count, Totals[i], nil, ReadMethod)
	}
}
func numToFile(i int) string {
	return "TestFile" + strconv.Itoa(i) + ".in"
}

func setupTestFiles() {
	rand.Seed(1234)
	Totals = make([]int64, len(TestFileNumbers))
	for fileNo, Top := range TestFileNumbers {
		temp := 0
		total := int64(0)
		filename := numToFile(Top)
		os.Remove(filename)
		file, _ := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		randTop := Top / 2
		if randTop <= 0 {
			randTop = 1
		}
		for i := 0; i < Top; i++ {
			temp++
			randInt := rand.Intn(100)
			total += int64(randInt)
			append := strconv.Itoa(randInt)
			append += " "

			if rand.Intn(randTop) < temp {
				temp = 0
				append += "\n"
			}
			file.WriteString(append)
		}
		file.Close()
		Totals[fileNo] = total
	}
}

func removeTestFiles() {
	for _, Top := range TestFileNumbers {
		filename := "TestFile" + strconv.Itoa(Top) + ".in"
		os.Remove(filename)
	}
}

func check2(i int, t *testing.T, got, got2 interface{}, goterr error, want, want2 interface{}, wanterr error, Area string) {
	check(i, t, got, goterr, want, wanterr, Area)
	if !(reflect.DeepEqual(got2, want2)) {
		t.Errorf("TEST: %+v\t%+v\nGOT2:\n%+v\nWANT2:\n%+v\n", i, Area, got2, want2)
	}
}

func check(i int, t *testing.T, got interface{}, goterr error, want interface{}, wanterr error, Area string) {
	switch {
	case goterr == nil && wanterr != nil:
		t.Errorf("TEST: %+v\t%+v\nGOTERR:\nnil\nWANTERR:\n%+v\n", i, Area, wanterr.Error())
		break
	case goterr != nil && wanterr == nil:
		t.Errorf("TEST: %+v\t%+v\nGOTERR:\n%+v\nWANTERR:\nnil\n", i, Area, goterr.Error())
		break
	case !reflect.DeepEqual(goterr, wanterr):
		t.Errorf("TEST: %+v\t%+v\nGOTERR:\n%+v\nWANTERR:\n%+v\n", i, Area, goterr.Error(), wanterr.Error())
		break
	case !reflect.DeepEqual(got, want):
		t.Errorf("TEST: %+v\t%+v\nGOT:\n%+v\nWANT:\n%+v\n", i, Area, got, want)
	}
}

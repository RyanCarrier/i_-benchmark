package methodTest

import (
	"os"
	"reflect"
	"testing"
)

type Test struct {
	Title string
	F     func(*os.File) ([]int, error)
}

var Tests = []Test{
	Test{
		Title: "ReadFromIoutilReadAll",
		F:     ReadFromIoutilReadAll,
	},
	Test{
		Title: "FromIoutil",
		F:     FromIoutil,
	},
	Test{
		Title: "FromBufioReadBytes",
		F:     FromBufioReadBytes,
	},
	Test{
		Title: "FromBufioFmtScan",
		F:     FromBufioFmtScan,
	},
	Test{
		Title: "FromBufioScanner",
		F:     FromBufioScanner,
	},
	Test{
		Title: "FromBufioReaderRead",
		F:     FromBufioReaderRead,
	},
}

type TestFile struct {
	Filename string
	Count    int
	Total    int
}

var Files = []TestFile{
	TestFile{
		Filename: "test1",
		Count:    7,
		Total:    7,
	},
	TestFile{
		Filename: "test2",
		Count:    3,
		Total:    4,
	},
	TestFile{
		Filename: "test3",
		Count:    11,
		Total:    150,
	},
}

func TestMethods(t *testing.T) {
	for i, file := range Files {
		for _, test := range Tests {
			f, _ := os.Open(file.Filename)
			defer f.Close()
			ints, err := test.F(f)
			Count, Total := Calc(ints)
			check2(i, t, Count, Total, err, file.Count, file.Total, nil, test.Title)
		}
	}
}

func check2(i int, t *testing.T, got, got2 interface{}, goterr error, want, want2 interface{}, wanterr error, Area string) {
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
	case !(reflect.DeepEqual(got, want)):
		t.Errorf("TEST: %+v\t%+v\nGOT:\n%+v\nWANT:\n%+v\n", i, Area, got, want)
	case !(reflect.DeepEqual(got2, want2)):
		t.Errorf("TEST: %+v\t%+v\nGOT2:\n%+v\nWANT2:\n%+v\n", i, Area, got2, want2)
	}
}

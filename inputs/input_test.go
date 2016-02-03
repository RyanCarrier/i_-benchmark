package inputs

import (
	"reflect"
	"testing"
)

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

func check(i int, t *testing.T, got interface{}, goterr error, want interface{}, wanterr error) {
	switch {
	case goterr == nil && wanterr != nil:
		t.Errorf("TEST: %+v\nGOTERR:\nnil\nWANTERR:\n%+v\n", i, wanterr.Error())
		break
	case goterr != nil && wanterr == nil:
		t.Errorf("TEST: %+v\nGOTERR:\n%+v\nWANTERR:\nnil\n", i, goterr.Error())
		break
	case !reflect.DeepEqual(goterr, wanterr):
		t.Errorf("TEST: %+v\nGOTERR:\n%+v\nWANTERR:\n%+v\n", i, goterr.Error(), wanterr.Error())
		break
	case !reflect.DeepEqual(got, want):
		t.Errorf("TEST: %+v\nGOT:\n%+v\nWANT:\n%+v\n", i, got, want)
	}
}

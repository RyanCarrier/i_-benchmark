package main

import (
	"testing"
)

func TestInputConvertAll(t *testing.T) {
	cases := []struct {
		in    string
		count int
		total int
	}{
		{
			in:    "1 2 3",
			count: 3,
			total: 6,
		},
		{
			in:    "100 2000 30000",
			count: 3,
			total: 32100,
		},
		{
			in:    "-1 1 1 1",
			count: 4,
			total: 3,
		},
	}
	methods := []string{"fmtscan", "scan", "splitstrconv"}
	for iteration, c := range cases {
		for _, method := range methods {
			cfg := NewCfg()
			parseMethod = method
			convertLine([]byte(c.in))
			switch {
			case total != c.total:
				t.Errorf("CASE %d\nTYPE %s\nTOTAL MISSMATCH\nGOT:\n\t\t%d\nWANT:\n\t\t%d",
					iteration, parseMethod, total, c.total)
			case count != c.count:
				t.Errorf("CASE %d\nTYPE %s\nCOUNT MISSMATCH\nGOT:\n\t\t%d\nWANT:\n\t\t%d",
					iteration, parseMethod, count, c.count)
			}
		}
	}

}

package inputs

import "testing"

var Methods = []string{"fmtscan", "scan", "splitstrconv"}

func TestInputEvaluateLine(t *testing.T) {
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
			total: 2,
		},
	}
	for _, method := range Methods {
		for i, c := range cases {
			cfg := NewCfg()
			cfg.ParseMethod = method
			err := cfg.EvaluateLine([]byte(c.in))
			check2(i, t, cfg.Count, cfg.Total, err, c.count, c.total, nil, method)
		}
	}
}

func TestInputEvaluateAll(t *testing.T) {
	cases := []struct {
		in    string
		count int
		total int
	}{
		{
			in: `1 2 3
      4 5 6
      7`,
			count: 7,
			total: 28,
		},
		{
			in: `100 2000 30000
      1
      1`,
			count: 5,
			total: 32102,
		},
		{
			in: `-1 1
      1
      1
      -3
      3`,
			count: 6,
			total: 2,
		},
	}
	for _, method := range Methods {
		for i, c := range cases {
			cfg := NewCfg()
			cfg.ParseMethod = method
			err := cfg.EvaluateAll([]byte(c.in))
			check2(i, t, cfg.Count, cfg.Total, err, c.count, c.total, nil, method)
		}
	}

}

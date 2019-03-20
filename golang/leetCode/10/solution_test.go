package main

import (
	"testing"
)

type question struct {
	p para
	a ans
}

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

// 单元测试 go test http_test.go http.go
func Test_Problem0010(t *testing.T) {
	qs := []question{
		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "aa",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "aa",
			},
			a: ans{
				one: false,
			},
		},

		question{
			p: para{
				one: "aa",
				two: "a*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aa",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "aaaaaaaab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*c",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "z*t*x*c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "c*a*b",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "abc",
				two: ".*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*b.*",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*...b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "b",
				two: ".*..*b",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		if q.a.one != isMatch(q.p.one, q.p.two) {
			t.Errorf("s = %v,p = %v", q.p.one, q.p.two)
		}
	}
}

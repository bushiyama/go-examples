package supermarketcoupon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	cases := map[string]struct {
		n       int
		p       []int
		wantSum int
	}{
		"empty": {
			n:       0,
			p:       []int{},
			wantSum: 0,
		},
		"1": {
			n:       3,
			p:       []int{200, 400, 300},
			wantSum: 700,
		},
		"2": {
			n:       4,
			p:       []int{9, 10, 11, 12},
			wantSum: 36,
		},
		"3": {
			n:       5,
			p:       []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000},
			wantSum: 4500000000,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ret := logic(tc.n, tc.p)
			assert.Equal(t, tc.wantSum, ret)
		})
	}
}

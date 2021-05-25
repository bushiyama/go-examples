package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	cases := map[string]struct {
		n    int
		s    string
		want string
	}{
		"empty": {
			n:    0,
			s:    "",
			want: "",
		},
		"1": {
			n:    3,
			s:    "aba",
			want: "ba",
		},
		"2": {
			n:    7,
			s:    "sptaast",
			want: "past",
		},
		"3": {
			n:    30,
			s:    "ryfoxchyvfmsewlwpoyvhdjkbvdjsa",
			want: "rxcfmelwpoyhkbvdjsa",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ret := logic(tc.n, tc.s)
			assert.Equal(t, ret, tc.want)
		})
	}
}

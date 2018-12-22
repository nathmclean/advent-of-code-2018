package main

import (
	"testing"
)

func TestPair(t *testing.T) {
	cases := []struct{
		name string
		a, b rune
		result bool
	}{
		{
			name: "matching pair",
			a: 'a',
			b: 'A',
			result: true,
		},
		{
			name: "another matching pair",
			a: 'A',
			b: 'a',
			result: true,
		},
		{
			name: "same case pair",
			a: 'a',
			b: 'a',
			result: false,
		},
		{
			name: "same case upper pair",
			a: 'A',
			b: 'A',
			result: false,
		},
		{
			name: "no matching same case pair",
			a: 'B',
			b: 'A',
			result: false,
		},
		{
			name: "no matching same case pair",
			a: 'b',
			b: 'a',
			result: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			result := pair(c.a, c.b)
			if result != c.result {
				t.Errorf("expected %v got %v", c.result, result)
			}
		})
	}
}

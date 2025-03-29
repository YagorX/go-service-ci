package main

import "testing"

func TestSum(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"sum three integers": {
			input: []int{1, 2, 3},
			want:  6,
		},
		"sum two integers": {
			input: []int{1, 2},
			want:  3,
		},
		"empty input": {
			input: []int{},
			want:  0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := Sum(tc.input); got != tc.want {
				t.Errorf("Sum() = %v, want %v", got, tc.want)
			}
		})
	}
}

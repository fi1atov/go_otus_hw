package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Переписал TestShapes с использованием tdt.
func TestWordCounterTdt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		row  string
		want map[string]int
	}{
		{
			desc: "case1",
			row:  "a h h a f a",
			want: map[string]int{"f": 1, "h": 2, "a": 3},
		},
		{
			desc: "case2",
			row:  "a h;h a,f a",
			want: map[string]int{"f": 1, "h": 2, "a": 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countWords(tC.row)
			assert.Equal(t, tC.want, got)
		})
	}
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounterTdt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		data []int
		find int
		want int
	}{
		{
			desc: "case1",
			data: []int{1, 5, 12, 18, 20, 26, 46, 78, 82},
			find: 12,
			want: 2,
		},
		{
			desc: "case2",
			data: []int{1, 5, 12, 18, 20, 26, 46, 78, 82},
			find: 78,
			want: 7,
		},
		{
			desc: "case3",
			data: []int{1, 5, 12, 18, 20, 26, 46, 78, 82},
			find: 1000,
			want: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := binarySearch(tC.data, tC.find)
			assert.Equal(t, tC.want, got)
		})
	}
}

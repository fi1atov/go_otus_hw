package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		want int
	}{
		{
			desc: "case1",
			want: 10,
		},
		{
			desc: "case2",
			want: 100,
		},
		{
			desc: "case3",
			want: 1000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := counterInGorutines(tC.want)
			assert.Equal(t, tC.want, got)
		})
	}
}

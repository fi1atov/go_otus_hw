package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritingToChannel(t *testing.T) {
	t.Parallel()
	data := [30]float64{
		21, 3, 57, 17, 211, 13, 8, 542, 45, 13,
		32, 357, 171, 94, 62, 2, 722,
		450, 50, 127, 27, 45, 99, 36, 25,
		348, 80, 549, 450, 11,
	}

	SensorDataChannel := make(chan float64, 10)

	go func() {
		for _, el := range data {
			// time.Sleep(1 * time.Second)
			SensorDataChannel <- el
			// fmt.Println("Insert to channel: ", el)
		}
	}()

	channel := SensorReader(SensorDataChannel)

	testCases := []struct {
		desc             string
		channelProcessed chan float64
		want             float64
	}{
		{
			desc:             "case1",
			channelProcessed: channel,
			want:             93.0,
		},
		{
			desc:             "case2",
			channelProcessed: channel,
			want:             206.7,
		},
		{
			desc:             "case3",
			channelProcessed: channel,
			want:             167.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := <-tC.channelProcessed
			assert.Equal(t, tC.want, got)
		})
	}

	// time.Sleep(2 * time.Second)	// это теперь не нужно
}

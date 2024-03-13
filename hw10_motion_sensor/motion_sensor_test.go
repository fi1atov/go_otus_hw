package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWritingToChannel(t *testing.T) {
	t.Parallel()
	data := [30]int{
		21, 3, 57, 17, 211, 13, 8, 542, 45, 13,
		32, 357, 171, 94, 62, 2, 722,
		450, 50, 127, 27, 45, 99, 36, 25,
		348, 80, 549, 450, 11,
	}

	SensorDataChannel := make(chan int, 10)
	ProcessedDataChannel := make(chan int)

	go func() {
		for _, el := range data {
			// time.Sleep(1 * time.Second)
			SensorDataChannel <- el
			fmt.Println("Insert to channel: ", el)
		}
	}()

	go processingData(SensorDataChannel, ProcessedDataChannel)

	testCases := []struct {
		desc             string
		channelProcessed chan int
		want             int
	}{
		{
			desc:             "case1",
			channelProcessed: ProcessedDataChannel,
			want:             93,
		},
		{
			desc:             "case2",
			channelProcessed: ProcessedDataChannel,
			want:             206,
		},
		{
			desc:             "case3",
			channelProcessed: ProcessedDataChannel,
			want:             167,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := <-tC.channelProcessed
			fmt.Println("got: ", got)
			fmt.Println("want: ", tC.want)
			assert.Equal(t, tC.want, got)
		})
	}

	time.Sleep(3 * time.Second)
}

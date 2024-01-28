package main

import (
	"testing"

	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	"github.com/stretchr/testify/assert"
)

func TestChessboardValue(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc           string
		chessboardSize int
		want           string
	}{
		{
			desc:           "four",
			chessboardSize: 4,
			want:           "# # \n # #\n# # \n # #\n",
		},
		{
			desc:           "six",
			chessboardSize: 6,
			want:           "# # # \n # # #\n# # # \n # # #\n# # # \n # # #\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := chessboard.GetChessBoard(tC.chessboardSize)
			assert.Equal(t, tC.want, got)
		})
	}
}

package main

import (
	"reflect"
	"testing"

	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	"github.com/stretchr/testify/assert"
)

var chessboardSize = 4

func TestChessboardType(t *testing.T) {
	t.Parallel()
	got := chessboard.GetChessBoard(chessboardSize)

	assert.Equal(t, reflect.TypeOf(got).Kind(), reflect.String)
}

func TestChessboardValue(t *testing.T) {
	t.Parallel()
	got := chessboard.GetChessBoard(chessboardSize)
	want := "# # \n # #\n# # \n # #\n"

	assert.Equal(t, want, got)
}

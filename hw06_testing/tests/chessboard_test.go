package tests

import (
	"reflect"
	"testing"

	chessboard "github.com/fi1atov/go_otus_hw/hw06_testing/hw03_chessboard"
	"github.com/stretchr/testify/assert"
)

func TestChessboard(t *testing.T) {
	t.Parallel()
	res := chessboard.GetChessBoard(10)
	assert.Equal(t, reflect.TypeOf(res).Kind(), reflect.String)
}

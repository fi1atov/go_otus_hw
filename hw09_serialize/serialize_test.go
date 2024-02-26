package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestStructJSON(t *testing.T) {
	t.Parallel()
	bk1 := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bk2 := Book{}

	got := bk1.MarshalJSON()
	want, _ := json.Marshal(&bk1)

	assert.Equal(t, want, got)

	bk2.UnmarshalJSON(got)

	assert.Equal(t, &bk1, &bk2)
}

func TestStructPROTO(t *testing.T) {
	t.Parallel()
	bk1 := Book{
		Isbn:   1,
		Title:  "Книга",
		Author: "Автор",
		Year:   1996,
		Size:   145,
		Rate:   0.8,
	}
	bk2 := Book{}

	got := bk1.MarshalPROTO()
	want, _ := proto.Marshal(&bk1)

	assert.Equal(t, want, got)

	bk2.UnmarshalPROTO(got)

	assert.Equal(t, &bk1, &bk2)
}

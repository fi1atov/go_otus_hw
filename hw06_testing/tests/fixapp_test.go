package tests

import (
	"reflect"
	"testing"

	jsondata "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app"
	"github.com/stretchr/testify/assert"
)

func TestFixApp(t *testing.T) {
	t.Parallel()
	res, _ := jsondata.GetJSONData()
	assert.Equal(t, reflect.TypeOf(res).Kind(), reflect.Slice)
}

package main

import (
	"reflect"
	"testing"

	jsondata "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app"
	employee "github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func getEmployees() (res []employee.Employee) {
	e1 := employee.Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3}
	e2 := employee.Employee{UserID: 11, Age: 30, Name: "George", DepartmentID: 2}

	res = append(res, e1)
	res = append(res, e2)

	return
}

func TestFixApp(t *testing.T) {
	t.Parallel()
	got, _ := jsondata.GetJSONData()
	want := getEmployees()
	assert.Equal(t, reflect.TypeOf(got).Kind(), reflect.Slice)
	assert.Equal(t, want, got)
}

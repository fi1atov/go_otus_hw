package jsondata

import (
	"fmt"

	"github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app/reader"
	"github.com/fi1atov/go_otus_hw/hw06_testing/hw02_fix_app/types"
)

func GetJSONData() (staff []types.Employee, err error) {
	path := "hw02_fix_app/data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	return staff, err
}

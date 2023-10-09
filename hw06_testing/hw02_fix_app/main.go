package main

import (
	"fmt"

	"github.com/ch1s7ya/otus_go_basic_hw/hw06_testing/hw02_fix_app/printer"
	"github.com/ch1s7ya/otus_go_basic_hw/hw06_testing/hw02_fix_app/reader"
	"github.com/ch1s7ya/otus_go_basic_hw/hw06_testing/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	printer.PrintStaff(staff)
}

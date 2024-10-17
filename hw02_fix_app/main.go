package main

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw02_fix_app/printer"
	"github.com/sar0868/otus_go_basic_hw/hw02_fix_app/reader"
	"github.com/sar0868/otus_go_basic_hw/hw02_fix_app/types"
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
		fmt.Printf("Error: %v", err)
	}

	printer.PrintStaff(staff)
}

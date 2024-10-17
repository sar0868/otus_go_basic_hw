package printer

import (
	"fmt"

	"github.com/sar0868/otus_go_basic_hw/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		pattern := "User ID: %d; Age: %d; Name: %s; Department ID: %d; "
		str = fmt.Sprintf(pattern, staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str)
	}

	fmt.Println(str)
}

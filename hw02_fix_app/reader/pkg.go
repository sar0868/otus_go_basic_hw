package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sar0868/otus_go_basic_hw/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}

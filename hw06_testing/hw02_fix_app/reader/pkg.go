package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ch1s7ya/otus_go_basic_hw/hw06_testing/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return data, nil
}

package internal

import (
	"os"
)

func ReadFiles(paths ...string) ([][]byte, error) {
	var data [][]byte

	for _, path := range paths {
		value, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	return data, nil
}
package mahoa

import (
	"io/ioutil"
)

func WriteToFile(filePath string, fileContent []byte) error {
	err := ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadBytesFromFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CombineBytes(a, b []byte) []byte {
	combined := make([]byte, len(a)+len(b))
	copy(combined[:len(a)], a)
	copy(combined[len(a):], b)
	return combined
}

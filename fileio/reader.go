package fileio

import (
	"os"

	"github.com/xaxys/bubbler/definition"
)

func GetFileContent(file *definition.FileIdentifer) (string, error) {
	content, err := ReadFile(file.Path)
	if err != nil {
		return "", &definition.GeneralError{
			Err: &definition.FileReadError{
				File: file,
				Err:  err,
			},
		}
	}
	return string(content), nil
}

// ReadFile returns unwarpped error
func ReadFile(path string) ([]byte, error) {
	inputFile, err := GetAbsolutelyPath(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	return content, nil
}

package fileio

import (
	"os"
	"path/filepath"

	"github.com/xaxys/bubbler/definition"
)

func WriteFileContent(file *definition.FileIdentifer, content string) error {
	err := WriteFile(file.Path, content)
	if err != nil {
		return &definition.FileWriteError{
			File: file,
			Err:  err,
		}
	}
	return nil
}

// WriteFile returns unwarpped error
func WriteFile(path string, content string) error {
	absPath, err := GetAbsolutelyPath(path)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(absPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(absPath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

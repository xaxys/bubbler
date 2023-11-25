package fileio

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/xaxys/bubbler/definition"
)

func GetFileIdentifer(file string) (*definition.FileIdentifer, error) {
	path, err := GetAbsolutelyPath(file)
	if err != nil {
		// unexpected error, return as is
		return nil, err
	}
	ident := &definition.FileIdentifer{
		Name: file,
		Path: path,
	}
	return ident, nil
}

func GetFileExistingStatus(file *definition.FileIdentifer) error {
	if err := CheckFileExist(file.Path); err != nil {
		return &definition.FileNotFoundError{
			File: file,
			Err:  err,
		}
	}
	return nil
}

// GetAbsolutelyPath returns unwarpped error
func GetAbsolutelyPath(file string) (absPath string, err error) {
	return GetAbsolutelyPathWithBasePath(file, "")
}

// GetAbsolutelyPathWithBasePath returns unwarpped error
func GetAbsolutelyPathWithBasePath(file string, relative string) (absPath string, err error) {
	if filepath.IsAbs(file) {
		absPath = file
	} else {
		absPath, err = filepath.Abs(filepath.Join(filepath.Dir(relative), file))
		if err != nil {
			return "", err
		}
	}
	return absPath, nil
}

// CheckFileExist returns unwarpped error
func CheckFileExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			return pathErr.Err
		}
		return err
	}
	return nil
}

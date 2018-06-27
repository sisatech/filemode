package filemode

import (
	"os"
)

// Create creates a new file at 'path' and sets permissions defined
// by 'mode'.
func Create(path string, mode os.FileMode) (*os.File, error) {

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	err = os.Chmod(path, mode)
	if err != nil {
		os.Remove(path)
		return nil, err
	}

	return f, nil
}

// Mkdir create a new directory at 'path' and sets permissions defined
// by 'mode'.
func Mkdir(path string, mode os.FileMode) error {

	err := os.Mkdir(path, mode)
	if err != nil {
		return err
	}

	// Go1.10 seems to disregard the 'mode' argument...
	err = os.Chmod(path, mode)
	if err != nil {
		os.Remove(path)
		return err
	}

	return nil
}

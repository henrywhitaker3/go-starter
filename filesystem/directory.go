package filesystem

import (
	"errors"
	"io"
	"os"
)

type Directory struct {
	Path string
}

// Create a new directory object
// This will create/assert the given path exists and has write perms
func NewDirectory(path string) (*Directory, error) {
	dir := Directory{Path: path}
	err := dir.createOrAssertPermissions()
	if err != nil {
		return nil, err
	}

	return &dir, nil
}

// Creates the new directory, or asserts that we have write permissions for it
func (d *Directory) createOrAssertPermissions() error {
	if d.Exists() {
		isEmpty, err := d.IsEmpty()
		if err != nil {
			return err
		}
		if !isEmpty {
			return errors.New("target directory must be empty")
		}
		return nil
	}

	if err := d.Create(); err != nil {
		return err
	}

	return nil
}

// Create the new directory
func (d Directory) Create() error {
	err := os.Mkdir(d.Path, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Determine if the directory already exists
func (d Directory) Exists() bool {
	_, err := os.Stat(d.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
	}

	return true
}

// Determine if the directory is empty or not
func (d Directory) IsEmpty() (bool, error) {
	f, err := os.Open(d.Path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// if the file is EOF the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

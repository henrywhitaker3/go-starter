package filesystem

import (
	"errors"
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
		return nil
	}

	if err := d.Create(); err != nil {
		return err
	}

	return nil
}

func (d Directory) Create() error {
	err := os.Mkdir(d.Path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (d Directory) Exists() bool {
	_, err := os.Stat(d.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
	}

	return true
}

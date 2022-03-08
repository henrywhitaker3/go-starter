package filesystem

import (
	"os"
)

type File struct {
	Name     string
	Contents string
}

// Copy the contents to a new file in a given directory
func (f File) CopyTo(dir Directory) error {
	err := os.WriteFile(dir.Path+"/"+f.Name, []byte(f.Contents), 0755)
	if err != nil {
		return err
	}

	return nil
}

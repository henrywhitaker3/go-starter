/*
Copyright © 2022 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package main

import (
	_ "embed"

	"github.com/henrywhitaker3/go-starter/cmd"
	"github.com/henrywhitaker3/go-starter/filesystem"
)

var (
	//go:embed stubs/.editorconfig
	editorconfig string
	//go:embed stubs/.gitignore
	gitignore string
	//go:embed stubs/Makefile
	makefile string
	//go:embed stubs/main.go
	maingo string
)

func main() {
	cmd.Files = []filesystem.File{
		{Name: ".editorconfig", Contents: editorconfig},
		{Name: ".gitignore", Contents: gitignore},
		{Name: "Makefile", Contents: makefile},
		{Name: "main.go", Contents: maingo},
	}

	cmd.Execute()
}

/*
Copyright © 2022 Henry Whitaker <henrywhitaker3@outlook.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/go-starter/filesystem"
	"github.com/spf13/cobra"
)

var (
	Files []filesystem.File
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gostarter [directory]",
	Short:   "A CLI to setup a default env for a go project.",
	Long:    `A CLI to setup a default env for a go project.`,
	Version: "0.1.0",
	Args:    cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		directoryArg := args[0]
		dir, err := filesystem.NewDirectory(directoryArg)
		if err != nil {
			return err
		}

		for _, file := range Files {
			if err := file.CopyTo(*dir); err != nil {
				return err
			}
		}

		fmt.Println("Open the new directory and run `go mod init [package name]`")

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-starter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

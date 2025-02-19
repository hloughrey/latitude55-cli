/*
Copyright © 2024 Hugh Loughrey
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

// TODO: Add a version command to the CLI
var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Latitude55 CLI",
  Long:  `All software has versions. This is Latitude55 CLI's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Latitude55 CLI v0.9 -- HEAD")
  },
}
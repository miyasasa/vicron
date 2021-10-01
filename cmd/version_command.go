package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints version of vicron",
		Long:  "Prints version of vicron",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("1.0.0")
			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(VersionCommand())
}

package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pdftool",
	Long:  `Prints the version number of pdftool`,
	//Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pdftool %s\n", VersionNumber)
	},
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lygg [file]",
	Short: "lygg – The log processing interpreter",
	Long:  "lygg – The log processing interpreter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(cmd.UsageString())
	},
}

func Execute() error {
	return rootCmd.Execute()
}

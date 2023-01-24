package cmd

import (
	"c2redir/cmd/httpcmd"

	"github.com/spf13/cobra"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
}

func init() {
	rootCmd.AddCommand(HttpCmd)

	HttpCmd.AddCommand(httpcmd.HttpInit)
}

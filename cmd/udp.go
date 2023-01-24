package cmd

import (
	"c2redir/cmd/udpcmd"

	"github.com/spf13/cobra"
)

var UdpCmd = &cobra.Command{
	Use:   "udp",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
}

func init() {
	rootCmd.AddCommand(UdpCmd)

	UdpCmd.AddCommand(udpcmd.UdpList)
	UdpCmd.AddCommand(udpcmd.UdpAdd)
	UdpCmd.AddCommand(udpcmd.UdpDel)
}

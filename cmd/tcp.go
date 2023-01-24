package cmd

import (
	"c2redir/cmd/tcpcmd"

	"github.com/spf13/cobra"
)

var TcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
}

func init() {
	rootCmd.AddCommand(TcpCmd)

	TcpCmd.AddCommand(tcpcmd.TcpList)
	TcpCmd.AddCommand(tcpcmd.TcpAdd)
	TcpCmd.AddCommand(tcpcmd.TcpDel)
}

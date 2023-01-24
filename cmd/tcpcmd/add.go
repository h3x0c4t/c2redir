package tcpcmd

import (
	"c2redir/utils/tcp"

	"github.com/spf13/cobra"
)

var TcpAdd = &cobra.Command{
	Use:   "add",
	Short: "add tcp connections",
	Run: func(cmd *cobra.Command, args []string) {
		tcp.AddTCP(cmd.Flag("ip").Value.String(), cmd.Flag("lport").Value.String(), cmd.Flag("rport").Value.String())
	},
}

func init() {
	TcpAdd.Flags().StringP("ip", "i", "", "IP address")
	TcpAdd.MarkFlagRequired("ip")
	TcpAdd.Flags().StringP("lport", "l", "", "Local port")
	TcpAdd.MarkFlagRequired("lport")
	TcpAdd.Flags().StringP("rport", "r", "", "Remote port")
	TcpAdd.MarkFlagRequired("rport")
}

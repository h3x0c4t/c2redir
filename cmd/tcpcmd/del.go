package tcpcmd

import (
	"c2redir/utils/tcp"

	"github.com/spf13/cobra"
)

var TcpDel = &cobra.Command{
	Use:   "del",
	Short: "del tcp connections",
	Run: func(cmd *cobra.Command, args []string) {
		tcp.DelTCP(cmd.Flag("ip").Value.String(), cmd.Flag("lport").Value.String(), cmd.Flag("rport").Value.String())
	},
}

func init() {
	TcpDel.Flags().StringP("ip", "i", "", "IP address")
	TcpDel.MarkFlagRequired("ip")
	TcpDel.Flags().StringP("lport", "l", "", "Local port")
	TcpDel.MarkFlagRequired("lport")
	TcpDel.Flags().StringP("rport", "r", "", "Remote port")
	TcpDel.MarkFlagRequired("rport")
}

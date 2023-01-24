package udpcmd

import (
	"c2redir/utils/udp"

	"github.com/spf13/cobra"
)

var UdpDel = &cobra.Command{
	Use:   "del",
	Short: "del udp connections",
	Run: func(cmd *cobra.Command, args []string) {
		udp.DelUDP(cmd.Flag("ip").Value.String(), cmd.Flag("lport").Value.String(), cmd.Flag("rport").Value.String())
	},
}

func init() {
	UdpDel.Flags().StringP("ip", "i", "", "IP address")
	UdpDel.MarkFlagRequired("ip")
	UdpDel.Flags().StringP("lport", "l", "", "Local port")
	UdpDel.MarkFlagRequired("lport")
	UdpDel.Flags().StringP("rport", "r", "", "Remote port")
	UdpDel.MarkFlagRequired("rport")
}

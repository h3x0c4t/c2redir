package udpcmd

import (
	"c2redir/utils/udp"

	"github.com/spf13/cobra"
)

var UdpAdd = &cobra.Command{
	Use:   "add",
	Short: "add udp connections",
	Run: func(cmd *cobra.Command, args []string) {
		udp.AddUDP(cmd.Flag("ip").Value.String(), cmd.Flag("lport").Value.String(), cmd.Flag("rport").Value.String())
	},
}

func init() {
	UdpAdd.Flags().StringP("ip", "i", "", "IP address")
	UdpAdd.MarkFlagRequired("ip")
	UdpAdd.Flags().StringP("lport", "l", "", "Local port")
	UdpAdd.MarkFlagRequired("lport")
	UdpAdd.Flags().StringP("rport", "r", "", "Remote port")
	UdpAdd.MarkFlagRequired("rport")
}

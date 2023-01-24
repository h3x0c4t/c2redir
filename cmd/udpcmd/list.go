package udpcmd

import (
	"c2redir/utils/udp"

	"github.com/spf13/cobra"
)

var UdpList = &cobra.Command{
	Use:   "list",
	Short: "list udp connections",
	Run: func(cmd *cobra.Command, args []string) {
		udp.ListUDP()
	},
}

func init() {
	// pass
}

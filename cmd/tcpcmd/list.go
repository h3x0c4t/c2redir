package tcpcmd

import (
	"c2redir/utils/tcp"

	"github.com/spf13/cobra"
)

var TcpList = &cobra.Command{
	Use:   "list",
	Short: "list tcp connections",
	Run: func(cmd *cobra.Command, args []string) {
		tcp.ListTCP()
	},
}

func init() {
	// pass
}

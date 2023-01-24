package httpcmd

import (
	"c2redir/utils/http"
	"log"

	"github.com/spf13/cobra"
)

var HttpInit = &cobra.Command{
	Use:   "init",
	Short: "init http redirection",
	Run: func(cmd *cobra.Command, args []string) {
		if http.EnableRewriteEngine() {
			http.EnableMod()
			http.RestartApache()
			http.AddHtaccess()
			log.Println("You can add your rules to /var/www/html/.htaccess")
		} else {
			log.Println("You already enabled rewrite engine, add your rules to /var/www/html/.htaccess")
		}
	},
}

func init() {
	// pass
}

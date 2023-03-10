package server

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	dirToServe string
	serverPort string
	urlPrefix  string
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "spools up a simple http file server",
	Long: `Quickly serve a local filesystem directory over http
With no arguments, http starts serving files under ./ over port 8080.
With '-d' arg, specify the directory to be served.
With '-p' arg, specify the port to serve on.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir(dirToServe))

		http.Handle(urlPrefix, http.StripPrefix(urlPrefix, fs))

		log.Printf("Serving %s over localhost:%s... Stop with ^C", dirToServe, serverPort)
		http.ListenAndServe(":"+serverPort, nil)
	},
}

func init() {
	ServerCmd.Flags().StringVarP(&dirToServe, "dir", "d", "./serve", "root directory to be served (ex: /var/www) [default is ./]")

	ServerCmd.Flags().StringVar(&serverPort, "port", "2020", "port to listen to [default is 2020)")
	ServerCmd.Flags().StringVar(&urlPrefix, "prefix", "/", "prefix required (ex: /static), suffix to host:port [default is /]")
	ServerCmd.AddCommand(httpCmd)
}

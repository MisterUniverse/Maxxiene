package server

import (
	"html/template"
	"net/http"
	"path"

	"github.com/spf13/cobra"
)

func httpTempCall() {
	dir := path.Join("index.html")
	html, err := template.ParseFiles(dir)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("html"))))

	http.HandleFunc(dir, func(rw http.ResponseWriter, req *http.Request) {

		if err != nil {
			http.Error(rw, "parse file error", 400)
			return
		}

		content := map[string]string{"name": "john doe"}

		err = html.Execute(rw, content)

		if err != nil {
			http.Error(rw, "render file error", 400)
			return
		}

	})

	http.ListenAndServe(":3000", nil)
}

// httpTempCmd represents the httpTemp command
var httpTempCmd = &cobra.Command{
	Use:   "httpTemp",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httpTempCall()
	},
}

func init() {
	ServerCmd.AddCommand(httpTempCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpTempCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpTempCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

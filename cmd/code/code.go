package code

import (
	"github.com/spf13/cobra"
)

type Template struct {
	templates map[string]string
}

var templatePaths = map[string]map[string]string{
	"html": {
		"t1": "./templates/html/t1.html",
		"t2": "./templates/html/t2.html",
		"t3": "./templates/html/t3.html",
		"t4": "./templates/html/t4.html",
	},
	"js": {
		"t1": "./templates/js/t1.js",
		"t2": "./templates/js/t2.js",
		"t3": "./templates/js/t3.js",
		"t4": "./templates/js/t4.js",
	},
	"css": {
		"t1": "./templates/css/t1.css",
		"t2": "./templates/css/t2.css",
		"t3": "./templates/css/t3.css",
		"t4": "./templates/css/t4.css",
	},
	"golang": {
		"t1": "./templates/go/t1.go",
		"t2": "./templates/go/t2.go",
		"t3": "./templates/go/t3.go",
		"t4": "./templates/go/t4.go",
	},
	"python": {
		"t1": "./templates/python/t1.py",
		"t2": "./templates/python/t2.py",
		"t3": "./templates/python/t3.py",
		"t4": "./templates/python/t4.py",
	},
}

func NewTemplates() *Template {
	templates := &Template{
		templates: make(map[string]string),
	}
	for lang, paths := range templatePaths {
		for name, path := range paths {
			key := lang + "_" + name
			templates.templates[key] = path
		}
	}
	return templates
}

var templates = NewTemplates()

// codeCmd represents the code command
var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "The 'code' pallette is used for setting up quick work environments",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}

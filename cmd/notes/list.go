/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package notes

import (
	"fmt"
	"log"
	"maxx/pkg/db"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var isNote bool

var dataType string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List different types of data",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("DATABASE"))

		var items []db.ItemScanner
		var err error

		switch dataType {
		case "photo":
			items, err = db.MaxxDB.Storage.ListItems("pictures", &db.Picture{})
		case "hexdmp":
			items, err = db.MaxxDB.Storage.ListItems("hex_dumps", &db.HexDump{})
		case "memdmp":
			items, err = db.MaxxDB.Storage.ListItems("memory_dumps", &db.MemoryDump{})
		case "file":
			items, err = db.MaxxDB.Storage.ListItems("files", &db.File{})
		case "note":
			items, err = db.MaxxDB.Storage.ListItems("notes", &db.Note{})
		default:
			log.Fatalf("Unknown data type: %s", dataType)
		}

		logError(err)

		for _, item := range items {
			printItem(item)
		}
	},
}

func printItem(item db.ItemScanner) {
	switch v := item.(type) {
	case *db.Picture:
		fmt.Printf("ID: %d, Filename: %s\n", v.ID, v.Filename)
	case *db.HexDump:
		fmt.Printf("ID: %d, Description: %s\n", v.ID, v.Description)
	case *db.MemoryDump:
		fmt.Printf("ID: %d, Filename: %s\n", v.ID, v.Description)
	case *db.File:
		fmt.Printf("ID: %d, Filename: %s\n", v.ID, v.Filename)
	case *db.Note:
		fmt.Printf("ID: %d, Content: %s\n", v.ID, v.Content)
	default:
		fmt.Println("Unknown item type")
	}
}

func logError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	localAppData := os.Getenv("LOCALAPPDATA") + "\\maxxiene"
	viper.SetConfigFile(localAppData + "\\config\\.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%s\n", err)
	}
	listCmd.Flags().StringVarP(&dataType, "type", "t", "note", "Data type to list (photo, hexdmp, memdmp, file, note)")
	NotesCmd.AddCommand(listCmd)
}

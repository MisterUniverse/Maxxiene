/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package notes

import (
	"encoding/hex"
	"fmt"
	"maxx/pkg/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var isPhoto, isHexDump, isMemoryDump, isFile bool

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save different types of data",
	Run: func(cmd *cobra.Command, args []string) {
		db.MaxxDB.Storage = db.NewDataStorage(viper.GetString("paths.DATABASE"))
		switch {
		case isPhoto:
			fmt.Println("Saving a photo...")
			db.MaxxDB.Storage.InsertData("pictures", "filename, data", args[0], args[1])
		case isHexDump:
			fmt.Println("Saving a hex dump...")
			db.MaxxDB.Storage.InsertData("hex_dumps", "description, data", args[0], hex.EncodeToString([]byte(args[1])))
		case isMemoryDump:
			fmt.Println("Saving a memory dump...")
			db.MaxxDB.Storage.InsertData("memory_dumps", "description, data", args[0], []byte(args[1]))
		case isFile:
			fmt.Println("Saving a file...")
			db.MaxxDB.Storage.InsertData("files", "filename, data", args[0], []byte(args[1]))
		default:
			fmt.Println("Saving text data...")
			db.MaxxDB.Storage.InsertData("notes", "title", "content", args[0], args[1])
		}
	},
}

func init() {
	saveCmd.Flags().BoolVarP(&isPhoto, "photo", "p", false, "Save as a photo")
	saveCmd.Flags().BoolVarP(&isHexDump, "hexdmp", "x", false, "Save as a hex dump")
	saveCmd.Flags().BoolVarP(&isMemoryDump, "memdmp", "m", false, "Save as a memory dump")
	saveCmd.Flags().BoolVarP(&isFile, "file", "f", false, "Save as a file")
	NotesCmd.AddCommand(saveCmd)
}

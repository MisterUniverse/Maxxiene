/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	fm "maxx/filemanager"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var (
	item string
)

func addItemToCheckList(item string) (string, error) {
	data := make(chan string)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go produce(data, item, &wg)

	go write(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		return "Successfully added item to checklist", nil
	} else {
		return "", errors.New("failed to add item to checklist")
	}

}

func produce(data chan string, s string, wg *sync.WaitGroup) {
	n := s
	data <- n
	wg.Done()
}

func write(data chan string, done chan bool) {
	f, err := os.OpenFile("todo.md", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for d := range data {
		fm.AppendFile(f, d)
	}
	done <- true
}

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "This command will add a todo item to your todo checklist.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := addItemToCheckList("- [] " + item + " " + strings.Join(args, " ")); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	todoCmd.Flags().StringVarP(&item, "add", "a", "", "Add a todo item")

	if err := todoCmd.MarkFlagRequired("add"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(todoCmd)
}

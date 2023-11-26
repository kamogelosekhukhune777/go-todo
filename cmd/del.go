package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "deletes your task",
	Long: `deletes your task from your Todo list. 
	you can delete one task at a time.
	index you provide must be within the range of you Todo List.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("To many argumnets only the first one will be executed")
		}

		intValue, err := strconv.Atoi(args[0])
		//check: idx > len(args)
		if err != nil { //proper error handling
			fmt.Println("input not valid")
		}

		if intValue > len(args) && intValue < 0 {
			fmt.Println("index out range of your todo list")
			return
		}

		Todos.Delete(intValue)

		err = Todos.Store(TodoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

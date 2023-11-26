package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "marks your task complete",
	Long: `marks your task on your todo list as complete.
	you can't complete two task at the same time, please one task at a time or else results in an error`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("to many arguments, one at a time please") //error handling
			return
		}

		intValue, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		Todos.Complete(intValue)

		err = Todos.Store(TodoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

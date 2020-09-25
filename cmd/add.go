package cmd

import(
  "fmt"
  "strings"
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a Task",
  Run: func(cmd *cobra.Command, args []string) {
    task := strings.Join(args," ")
    fmt.Printf("Added Task %s to the List.\n",task)
  },
}

func init() {
  RootCmd.AddCommand(addCmd)
}

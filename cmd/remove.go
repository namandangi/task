package cmd

import(
  "fmt"
  "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
  Use:   "remove",
  Short: "Remove a Task",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Remove a Task")
  },
}

func init() {
  RootCmd.AddCommand(removeCmd)
}

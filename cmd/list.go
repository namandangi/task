package cmd

import(
  "fmt"
  "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List all Tasks",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("List Tasks")
  },
}

func init() {
  RootCmd.AddCommand(listCmd)
}

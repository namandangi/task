package cmd

import(
  "fmt"
  "os"
  "strings"
  "github.com/task/db"
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a Task",
  Run: func(cmd *cobra.Command, args []string) {
    task := strings.Join(args," ")
    _, err := db.CreateTask(task)
    if err != nil {
      fmt.Println("Something went wrong",err.Error())
      os.Exit(1)
    }
    fmt.Printf("Added Task %s to the List.\n",task)
  },
}

func init() {
  RootCmd.AddCommand(addCmd)
}

package cmd

import(
  "fmt"
  "strconv"
  "github.com/spf13/cobra"
  )

var doCmd = &cobra.Command{
  Use: "do",
  Short: "Do a Task",
  Run: func(cmd *cobra.Command, args []string) {
    var ids []int
    for _, arg := range args{
      id, err := strconv.Atoi(arg)
      if err == nil {
        ids = append(ids,id)
      } else {
        fmt.Println("Failed to parse the argument :",arg)
      }
    }
    fmt.Println(ids)
  },
}

func init() {
  RootCmd.AddCommand(doCmd)
}

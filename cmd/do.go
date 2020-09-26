package cmd

import(
  "fmt"
  "os"
  "strconv"
  "github.com/spf13/cobra"
  "github.com/task/db"
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
    tasks, err := db.AllTasks()
    if err != nil {
      fmt.Println("Something went wrong",err)
      os.Exit(1)
    }
    for _, id := range ids {
      if id <= 0 || id > len(tasks) {
        fmt.Println("Invalid task id",id)
        continue
      }
      task := tasks[id-1]
      err := db.DeleteTask(task.Key)
      if err != nil {
        fmt.Printf("Failed to mark Task %d as completed. Error: %s\n" ,id,err)
      } else {
        fmt.Printf("Marked Task %d as completed.\n",id)
      }
    }
  },
}

func init() {
  RootCmd.AddCommand(doCmd)
}

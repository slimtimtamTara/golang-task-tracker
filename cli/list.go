package cli

import (
	"github.com/spf13/cobra"
	"task-tracker/types"
	helpers "task-tracker/helpers"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"encoding/json"
	"fmt"
)

func init(){
	RootCmd.AddCommand(listCommand)
}
var listCommand = &cobra.Command{
	Use: "list",
	Short: "provides a list of all your tasks, completed and uncompleted",
	Run: func(cmd *cobra.Command, args []string){
		todos := &types.Tasks{}
		List(todos)
	}, 
}

func List(list *types.Tasks) error {
	data, err := os.ReadFile(helpers.GetTodoFile())
	helpers.HandleError(err)
	if len(data) == 0 {
		return err
	}
	err = json.Unmarshal(data, list)
	fmt.Println("Current Task  List:")
	printData(list)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func printData(list *types.Tasks){
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"ID", "Task", "Done", "Category","Created On", "Completed On"})
	for _, v := range *list {
		t.AppendRows([]table.Row{
			{v.ID,v.Task,v.Done,v.Category,v.CreatedOn,v.CompletedOn},
		})
		t.AppendSeparator()
	}
    t.AppendFooter(table.Row{"Total", len(*list)})
    t.Render()
}
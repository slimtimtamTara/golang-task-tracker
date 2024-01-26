package cli

import (
	"github.com/spf13/cobra"
	"task-tracker/types"
	helpers "task-tracker/helpers"
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
	fmt.Println(list)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
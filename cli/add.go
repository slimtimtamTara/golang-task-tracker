package cli

import (
	"github.com/spf13/cobra"
	helpers "task-tracker/helpers"
	"fmt"
	"task-tracker/types"
	"time"
	"github.com/google/uuid"
	"encoding/json"
	"os"
	"golang.org/x/exp/slices"
)

func init(){
	addCommand.PersistentFlags().String("cat", "c", "task category")
	RootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add [TASK]",
	Short: "adds a todo task to your list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		todos := &types.Tasks{}
		categoryOptions := []string{"work", "school", "fun"}
		category, _ := cmd.Flags().GetString("cat")
		task := args[0]
		if category != ""{
			if slices.Contains(categoryOptions, category){
				Add(task, category, todos)
				fmt.Printf("Successfully added %s\n", task)
			} else {
				fmt.Println("that category doesn't exist! The options are: work, school, or fun")
			}
		} else {
			fmt.Println("Give your task a category! The options are: work, school, or fun")
		}
	},
}

func Add(t string, cat string, list *types.Tasks) error {
	task := types.TaskItem{
		ID: uuid.New(),
		Task: t,
		Done: false,
		Category: cat,
		CreatedOn: time.Now(),
		CompletedOn: nil,
	}
	existingList, err := helpers.ReadTodosFromFile()
	helpers.HandleError(err)
	*list = append(existingList, task)
	fmt.Println(*list)


	data, err := json.Marshal(*list)
	helpers.HandleError(err)
	return os.WriteFile(helpers.GetTodoFile(), data, 0644)
	
}


package cli

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"path/filepath"
	helpers "task-tracker/helpers"
)

func init(){
	RootCmd.AddCommand(initCommand)
}

// SearchByFile Command
var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize local todo list",
	Long:  `This command creates a local task tracking json file for you`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		helpers.HandleError(err)
		filepath := filepath.Join(homeDir+"/Desktop/", ".todos.json")
		_, err = os.Stat(filepath)
		if err != nil {
			if os.IsNotExist(err) {
				file, err := os.Create(filepath)
				helpers.HandleError(err)
				defer file.Close()
				fmt.Println("created a todo json file on your Desktop ")
			}
		} else {
			fmt.Println("a todo file aleady exists")
		}
	},
}
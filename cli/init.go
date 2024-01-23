package cli

import (
	"github.com/spf13/cobra"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		if err != nil {
			log.Fatal(err)
		}
		filepath := filepath.Join(homeDir+"/Desktop/", ".todos.json")
		_, err = os.Stat(filepath)
		if err != nil {
			if os.IsNotExist(err) {
				file, err := os.Create(filepath)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				fmt.Println("created a todo json file on your Desktop ")
			}
		} else {
			fmt.Println("a todo file aleady exists")
		}
	},
}
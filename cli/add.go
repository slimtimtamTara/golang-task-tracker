package cli

import (
	"github.com/spf13/cobra"
)

func init(){
	RootCmd.AddCommand(addCommand)
}

// SearchByFile Command
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "add Task",
	Long:  `adds a todo task to your list`,
}


package cli

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "track-task",
	Short: "keep track of your tasks",
	Long:  `first time making a cli. I'm assuming this one is the outer layer?`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
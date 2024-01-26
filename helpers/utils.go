package helpers

import (
	"os"
	"path/filepath"
	"log"
	"encoding/json"
	"task-tracker/types"

)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func GetTodoFile() string {
	homeDir, err := os.UserHomeDir()
	HandleError(err)
	return filepath.Join(homeDir+"/Desktop/", ".todos.json")
}

func ReadTodosFromFile() (types.Tasks, error) {
	file, err := os.ReadFile(GetTodoFile())
	HandleError(err)
	var todos types.Tasks
	err = json.Unmarshal(file, &todos)
	HandleError(err)
	return todos, nil
}
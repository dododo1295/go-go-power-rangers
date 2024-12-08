package main

// Design Principles:
// Needs to have a way of storing and saving data into the system Documents folder (name: TBD)
// Needs to have commands to add, remove, rename, mark as complete, list all commands, have a nifty way of crossing out in the csv without deleting for completion (maybe moving to a different file????)
// Needs to have a quick and easy terminal command to be used for it
// Finally needs to stay within scope but also do a good job (would like to use this for myself)
// I also would like for this to be potentially scalable into a standalone app.

import (
	"fmt"
)

func main() {
	fmt.Println("This is the main.go file")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	storage.Save(todos)
}

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	a string
	d int
	e string
	t int
	l bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.a, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.e, "Edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.d, "del", -1, "Specify by index to delete")
	flag.IntVar(&cf.t, "toggle", -1, "Specify by index to toggle")
	flag.BoolVar(&cf.l, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.l:
		todos.print()
	case cf.a != "":
		todos.add(cf.a)
	case cf.e != "":
		parts := strings.SplitN(cf.e, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format, use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	case cf.t != -1:
		todos.toggle(cf.t)
	case cf.d != -1:
		todos.delete(cf.d)
	default:
		fmt.Println("invalid command")
	}
}

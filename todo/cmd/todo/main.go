package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const TODOFILENAME = "todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(TODOFILENAME); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(TODOFILENAME); err != nil {
			fmt.Fprintf(os.Stderr, "saving file error")
			os.Exit(1)
		}
	}
}

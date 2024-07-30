package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	gotodo "github.com/SyedAsadK/go-todo"
)

const (
	todofile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new task")
	done := flag.Int("done", 0, "mark a task as completed")
	del := flag.Int("del", 0, "deletes a task")
	list := flag.Bool("list", false, "List all tasks")
  

	flag.Parse()

	Todo := &gotodo.Todos{}

	err := Todo.Load(todofile)
  Err(err)	

	switch {

	case *add:
    task,err := getInput(os.Stdin,flag.Args()...)
    Err(err)
  
    Todo.Add(task)
		err = Todo.Store(todofile)
    Err(err)
	


	case *done> 0:
		err := Todo.Compelete(*done)
    Err(err)
		
		err = Todo.Store(todofile)
    Err(err)
		
	case *del> 0:
		err := Todo.Delete(*del)
    Err(err)
		
		err = Todo.Store(todofile)
    Err(err)
		
	case *list:
		Todo.Print()

	default:
		fmt.Fprintln(os.Stdout, "invalid command, type -h to see all commands")
		os.Exit(1)

	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
   
	scanner := bufio.NewScanner(r)
	scanner.Scan()
   
	if err := scanner.Err(); err != nil {
		return "", err
	}
  text := scanner.Text()

  if len(text) == 0 {
    return "",errors.New("Empty task")
  }
  return text,nil

}
func Err(err error){
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
}

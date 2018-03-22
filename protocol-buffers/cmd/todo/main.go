// We will make a very simple CLI tool

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/prantoran/justforfunc-prac/protocol-buffers/todo"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("TODO")
}

const dbPath = "mydb.pb"

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}
	fmt.Println(proto.MarshalTextString(task))
	return nil
}

func list() error {
	return nil
}

package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func init() {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Printf(" Exiting with ")
		os.Exit(1)
	}()
}

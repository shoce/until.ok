/*

go mod init github.com/shoce/until.ok
go get -a -u -v
go mod tidy

GoFmt
GoBuildNull
GoBuild

*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	NL = "\n"
)

func main() {
	var err error
	var Duration time.Duration
	var Command *exec.Cmd

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: until.ok duration command"+NL)
		os.Exit(1)
	}

	Duration, err = time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "time.ParseDuration %s: %v"+NL, os.Args[1], err)
		os.Exit(1)
	}

	for {
		Command = exec.Command(os.Args[2], os.Args[3:]...)
		Command.Stdin, Command.Stdout, Command.Stderr = os.Stdin, os.Stdout, os.Stderr
		fmt.Fprintf(os.Stderr, NL+"%s:"+NL, Command)
		err = Command.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, NL+"err: %v"+NL, err)
		}
		if err == nil {
			os.Exit(0)
		}

		fmt.Fprintf(os.Stderr, NL+"sleeping %v"+NL, Duration)
		time.Sleep(Duration)
	}

}

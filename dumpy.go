package main

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	line = "---------------------------------------------"
)

func main() {
	currenttime := time.Now()
	fmt.Printf("Dump started at: %s\n", currenttime.Format(time.RFC3339))
	// Commands array defined in seperate file (commands.go)
	runCommands(Commands)
}

// Iterate over our array of commands and print the results, padded with some titles and lines
func runCommands(commands []Command) {
	var out []byte
	var err error
	for _, c := range Commands {
		if len(c.Args) < 1 {
			out, err = exec.Command(c.Name).Output()
		} else {
			out, err = exec.Command(c.Name, c.Args).Output()
		}
		if err != nil {
			fmt.Println("[ERROR]: ", err)
		}

		fmt.Println(line)
		fmt.Printf("# [%s]:\n\n", c.Desc)
		fmt.Println(string(out))
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	changeOrigin := flag.String("c", "false", "commits to two repositories")
	commitMessage := flag.String("m", "World", "name to greet")
	flag.Parse()

	if *changeOrigin == "false" {
		commitPush(*commitMessage)
		fmt.Println("Pushed!")
	}
}

func commitPush(commitMessage string) {
	runCommand("git", "add", "*")
	runCommand("git", "commit", "-m", commitMessage)
	runCommand("git", "push")
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %s %v\nError: %v\n", name, args, err)
	} else {
		fmt.Printf("Successfully ran: %s %v\n", name, args)
	}
}

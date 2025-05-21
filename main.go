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
	} else {
		commitPushTwoRepos(*commitMessage)
		fmt.Println("Pushed to two repositories!")
	}
}

func commitPushTwoRepos(commitMessage string) {
	// Commit setup
	runCommand("git", "add", "*")
	runCommand("git", "commit", "-m", commitMessage)

	// Push to main repo
	runCommand("git", "remote", "add", "origin", os.Getenv("originalRepo"))
	runCommand("git", "push", "--set-upstream", "origin", "main")

	// Push to side repo
	runCommand("git", "remote", "rm", "origin")
	runCommand("git", "remote", "add", "origin", os.Getenv("sideRepo"))
	runCommand("git", "push", "--set-upstream", "origin", "main")

	// Clean up
	runCommand("git", "remote", "rm", "origin")
	runCommand("git", "remote", "add", "origin", os.Getenv("originalRepo"))
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

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

		originalRepo := "https://github.com/orladigital/indexa-web.git"
		sideRepo := "https://github.com/nandowastaken/indexa-web-vercel.git"

		commitPushTwoRepos(*commitMessage, originalRepo, sideRepo)
		fmt.Println("Pushed to two repositories!")
	}
}

func commitPushTwoRepos(commitMessage, originalRepo, sideRepo string) {
	// Commit setup
	runCommand("git", "add", "*")
	runCommand("git", "commit", "-m", commitMessage)

	// Push to main repo
	runCommand("git", "remote", "add", "origin", originalRepo)
	runCommand("git", "push", "--set-upstream", "origin", "main")

	// Push to side repo
	runCommand("git", "remote", "rm", "origin")
	runCommand("git", "remote", "add", "origin", sideRepo)
	runCommand("git", "push", "--set-upstream", "origin", "main")

	// Clean up
	runCommand("git", "remote", "rm", "origin")
	runCommand("git", "remote", "add", "origin", originalRepo)
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

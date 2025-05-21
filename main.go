package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	green := "\033[32m"

	changeOrigin := flag.String("c", "false", "commits to two repositories")
	commitMessage := flag.String("m", "World", "name to greet")
	flag.Parse()

	if *changeOrigin == "false" {
		commitPush(*commitMessage)
	} else {

	}

	fmt.Println(green, "Pushed!")
}

func commitPushTwoRepos(commitMessage string) {
	commitPush(commitMessage)
	exec.Command("git", "remote", "rm", "origin").Run()
	exec.Command("git", "remote", "add", "origin").Run()

	commitPush(commitMessage)
	exec.Command("git", "remote", "rm", "origin").Run()
	exec.Command("git", "remote", "add", "origin").Run()
}

func commitPush(commitMessage string) {
	exec.Command("git", "add", "*").Run()
	exec.Command("git", "commit", "-m", commitMessage).Run()
	exec.Command("git", "push").Run()
}

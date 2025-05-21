package main

import (
	"flag"
	"os/exec"
)

func main() {
	commitMessage := flag.String("m", "World", "name to greet")
	flag.Parse()
	// Yo

	exec.Command("git", "add", "*").Run()
	exec.Command("git", "commit", "-m", *commitMessage).Run()
	exec.Command("git", "push").Run()

}

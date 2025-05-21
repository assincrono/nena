package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	green := "\033[32m"
	commitMessage := flag.String("m", "World", "name to greet")
	flag.Parse()

	exec.Command("git", "add", "*").Run()
	exec.Command("git", "commit", "-m", *commitMessage).Run()
	exec.Command("git", "push").Run()

	fmt.Println(green, "Pushed!")
}

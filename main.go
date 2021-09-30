package main

import (
	"cmd2img/lib"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if 2 > len(os.Args) {
		fmt.Println("cmd2img <command> <output image name>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]
	shell := os.Getenv("SHELL")

	result, _ := exec.Command(shell, "-c", command).CombinedOutput()

	outputText := fmt.Sprintf("~$ %s\n%s", command, string(result))

	lib.DrawImage(outputText, filename)
}

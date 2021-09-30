package main

import (
	"cmd2img/lib"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	if 2 > len(os.Args) {
		fmt.Println("cmd2img <command> <output image name>")
		os.Exit(1)
	}

	args := os.Args[1 : len(os.Args)-1]
	filename := os.Args[len(os.Args)-1]
	shell := os.Getenv("SHELL")

	command := strings.Join(args, " ")
	result, _ := exec.Command(shell, "-c", command).CombinedOutput()

	outputText := fmt.Sprintf("~$ %s\n%s", command, string(result))

	lib.DrawImage(outputText, filename)
}

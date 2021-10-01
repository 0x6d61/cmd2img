package main

import (
	"cmd2img/lib"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var file = flag.String("f", "output.png", "image filename")
	flag.Parse()

	args := flag.Args()
	help := func() {
		fmt.Fprintf(os.Stderr, `
cmd2img [command] [option]

Usage:
   ~$ cmd2img ls -la
   ~$ cmd2img ls -la -f test.png
   ~$ cmd2img "ls | grep cmd2img"
	
   -f image filename default:output.png
	`)
	}
	flag.Usage = help
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	shell := os.Getenv("SHELL")

	command := strings.Join(args, " ")
	result, _ := exec.Command(shell, "-c", command).CombinedOutput()

	outputText := fmt.Sprintf("~$ %s\n%s", command, string(result))

	lib.DrawImage(outputText, *file)
}

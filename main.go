package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

func runCommand(command string, args ...string) string {
	commandBin, err := safeexec.LookPath(command)
	if err != nil {
		log.Fatalf("Error finding command %s, error %v\n", command, err)
	}

	cmd := exec.Command(commandBin, args...)
	cmd_output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error runnning command %s, error %v\n", command, err)
	}

	return strings.TrimSpace(string(cmd_output))
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: tsc-node <file-name>\n")
		fmt.Fprintf(os.Stderr, "Error: file name not provided\n")
	} else if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Usage: tsc-node <file-name>\n")
		fmt.Fprintf(os.Stderr, "Error: extra arguments provided\n")
	} else if _, err := os.Stat(args[0]); err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Error: file was not found\n")
	} else {
		tsFileName := args[0]
		jsFileName := strings.TrimSuffix(tsFileName, ".ts") + ".js"

		_ = runCommand("tsc", tsFileName)

		nodeCmdOut := runCommand("node", jsFileName)
		fmt.Println(nodeCmdOut)
	}
}

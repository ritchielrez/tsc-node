package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

const progName = "tsc-node"

func runCommand(command string, args ...string) string {
	commandBin, err := safeexec.LookPath(command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding command %s, error %v\n", command, err)
	}

	cmd := exec.Command(commandBin, args...)
	cmd_output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error runnning command %s, error %v\n", command, err)
	}

	return strings.TrimSpace(string(cmd_output))
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file-name>\n", progName)
		fmt.Fprintf(os.Stderr, "Error: file name not provided\n")
	} else if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file-name>\n", progName)
		fmt.Fprintf(os.Stderr, "Error: extra arguments provided\n")
	} else if _, err := os.Stat(args[0]); err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Error: file was not found\n")
	} else {
		tsFileName := args[0]
		jsFileName := strings.TrimSuffix(tsFileName, ".ts") + ".js"

		_, err := os.Stat(jsFileName)
        // If the js file is not found
		if err != nil && errors.Is(err, os.ErrNotExist) {
			tscCmdOut := runCommand("tsc", "--noEmitOnError", tsFileName)
			fmt.Println(tscCmdOut)
		}

        // If the js file is found
        // This check is done again because tsc may not generate js file,
        // if the ts file contains any error.
        if !errors.Is(err, os.ErrNotExist) {
            nodeCmdOut := runCommand("node", jsFileName)
            fmt.Println(nodeCmdOut)
        }
	}
}

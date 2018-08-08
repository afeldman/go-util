package print

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PrintCommand(cmd *exec.Cmd) {
	if debug {
		fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
	}
}

func PrintError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func PrintOutput(outs []byte) {
	if (len(outs) > 0) && (debug) {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

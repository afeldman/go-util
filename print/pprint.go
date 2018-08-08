package print

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var DEBUG = false

func PrintCommand(cmd *exec.Cmd) {
	if DEBUG {
		fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
	}
}

func PrintError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func PrintOutput(outs []byte) {
	if (len(outs) > 0) && DEBUG {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

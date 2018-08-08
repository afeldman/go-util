// +build !debug
package print

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PrintCommand(cmd *exec.Cmd) {
}

func PrintError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func PrintOutput(outs []byte) {
}

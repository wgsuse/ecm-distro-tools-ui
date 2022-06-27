package include

import (
	"os/exec"
	"runtime"
	"strings"
)

var toolList []string = nil

func Tools(path string) []string {
	if toolList != nil {
		return toolList
	}
	args := []string{"ls-files", "-o", "-x", "*.exe", "cmd/**/bin/*"}
	if runtime.GOOS == "windows" {
		args = []string{"ls-files", "-o", "cmd/**/*.exe"}
	}
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = SyscallAttrs()
	}

	if b, err := cmd.Output(); err == nil {
		toolList = strings.Split(string(b), "\n")

		return toolList
	}
	return []string{}
}

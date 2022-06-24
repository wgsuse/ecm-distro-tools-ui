package include

import (
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

var authorList []string = nil

func Authors() []string {
	if authorList != nil {
		return authorList
	}
	cmd := exec.Command("git", "log")
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = SyscallAttrs()
	}
	if b, err := cmd.Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Werner Garcia <werner.garcia@suse.com>"}
}

//go:build !windows
// +build !windows

package include

import "syscall"

func SyscallAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

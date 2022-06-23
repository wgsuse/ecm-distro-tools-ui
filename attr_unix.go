//go:build !windows
// +build !windows

package main

import "syscall"

func SyscallAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

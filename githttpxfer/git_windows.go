// +build windows

package githttpxfer

import (
	"os/exec"
	"syscall"
)

func (g *git) GitCommand(repoPath string, args ...string) *exec.Cmd {
	command := exec.Command(g.binPath, args...)
	command.Dir = g.GetAbsolutePath(repoPath)
	command.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	return command
}

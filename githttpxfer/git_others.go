// +build linux freebsd darwin

package githttpxfer

import (
	"os/exec"
	"syscall"
)

func (g *git) GitCommand(repoPath string, args ...string) *exec.Cmd {
	command := exec.Command(g.binPath, args...)
	command.Dir = g.GetAbsolutePath(repoPath)
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return command
}

package githttpxfer

import (
	"net/http"
	"os/exec"
	"strings"
)

func getServiceType(req *http.Request) string {
	serviceType := req.FormValue("service")
	if has := strings.HasPrefix(serviceType, "git-"); !has {
		return ""
	}
	return strings.Replace(serviceType, "git-", "", 1)
}

func cleanUpProcessGroup(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}
	if process := cmd.Process; process != nil && process.Pid > 0 {
		process.Kill()
	}
}

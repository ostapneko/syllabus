package syllabus

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DomainExists(domain string) bool {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "whois"
	cmdArgs := []string{domain}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error calling whois: %s", err))
		return false
	}

	if strings.HasPrefix(strings.TrimSpace(string(cmdOut)), "Domain Name") {
		return true
	}

	os.Stderr.WriteString(string(cmdOut))
	return false
}

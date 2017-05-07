package procedures

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"vmklet/model"
)

/*
* SSH Command stream
* SSH commands executor implements model.CommandStream
*/

type SSHCommandStream struct {
	User         string
	IP           string
	Port         int
	IdentityFile string
	Arguments    []string
}

func (ssh *SSHCommandStream) Command(cmd model.Command) error {
	return  nil
}

func (ssh *SSHCommandStream) Procedure(cmd ...model.Command) error {
	return  nil
}


func (ssh *SSHCommandStream) executeOne(cmd ...string) *exec.Cmd {
	var args []string = make([]string, 0)
	if ssh.IdentityFile != "" {
		args = append(args, "-i", ssh.IdentityFile)
	}

	if len(ssh.Arguments) > 0 {
		args = append(args, ssh.Arguments...)
	}

	args = append(args, fmt.Sprintf("%s@%s:%d", ssh.User, ssh.IP, ssh.Port))

	args = append(args, cmd...)

	return exec.Command("ssh", args...)
}

func (ssh *SSHCommandStream) executeMany(commands ...string) ([]string, error) {
	var err error
	var out []string = make([]string, 0)
	var args []string = make([]string, 0)
	if ssh.IdentityFile != "" {
		args = append(args, "-i", ssh.IdentityFile)
	}

	if len(ssh.Arguments) > 0 {
		args = append(args, ssh.Arguments...)
	}

	args = append(args, fmt.Sprintf("%s@%s:%d", ssh.User, ssh.IP, ssh.Port))

	cmd := exec.Command("ssh", args...)
	reader, err := cmd.StdoutPipe()
	scanner := bufio.NewScanner(reader)
	defer reader.Close()
	err = scanner.Err()
	if err == nil {
		for _, command := range commands {
			cmd.Stdin = strings.NewReader(command)
			for scanner.Scan() {
				err = scanner.Err()
				if err != nil {
					break
				}
				out = append(out, string(scanner.Bytes()))
			}
		}
	}

	// Graceful ssh exit
	cmd.Stdin = strings.NewReader("exit")

	time.Sleep(2 * time.Second)

	if cmd.Process.Pid > 0 {
		// Forced ssh exit
		cmd.Process.Kill()
	}

	return out, err
}

package exec

import (
	"bytes"
	"errors"
	"github.com/creack/pty"
	"io"
	"log/slog"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func New(opts ...Option) (*Command, error) {
	var err error
	c := &Command{
		cmd: exec.Command("", ""),
	}

	for _, opt := range opts {
		err = opt(c)
		if err != nil {
			return nil, err
		}
	}
	cmd = c.cmd
	slog.Debug("run command", "args", c.cmd.Args)

	if c.cmd.Cancel == nil {
		slog.Debug("run command kill global cmd")
		// 没有配置超时时间，则先杀掉上一个进程
		if cmd != nil && cmd.Process != nil && cmd.Process.Pid > 0 {
			if err := cmd.Process.Kill(); err == nil {
				_, err = cmd.Process.Wait()
				if err != nil {
					slog.Debug("run command kill global cmd", "error", err.Error())
				}
			}
		}
	}

	return c, nil
}

type Command struct {
	cmd *exec.Cmd
}

func (self Command) RunInTerminal(size *pty.Winsize) (io.ReadCloser, error) {
	var out *os.File
	var err error

	out, err = pty.StartWithSize(self.cmd, size)
	if err != nil {
		return nil, err
	}
	return TerminalResult{
		Conn: out,
		cmd:  self.cmd,
	}, err
}

func (self Command) Run() (io.Reader, error) {
	out := new(bytes.Buffer)
	self.cmd.Stdout = out
	self.cmd.Stderr = out
	err := self.cmd.Run()
	if err != nil {
		return nil, errors.Join(err, errors.New(out.String()))
	}
	return out, nil
}

func (self Command) RunWithResult() string {
	out, err := self.cmd.CombinedOutput()
	if err != nil {
		slog.Debug("run command with result", "arg", self.cmd.Args, "error", err.Error())
		return ""
	}
	return string(out)
}

func (self Command) Kill() error {
	if cmd != nil {
		err := cmd.Process.Kill()
		if err != nil {
			return err
		}
		return cmd.Wait()
	}
	return nil
}

func (self Command) Cmd() *exec.Cmd {
	return self.cmd
}

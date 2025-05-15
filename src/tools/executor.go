package tools

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CommandExecutor 用于执行命令并捕获输出
type CommandExecutor struct {
	Command string
	Args    []string
	Stdout  bytes.Buffer
	Stderr  bytes.Buffer
}

// NewCommandExecutor 创建一个新的 CommandExecutor 实例
func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{}
}

// Run 执行命令并等待其完成
func (ce *CommandExecutor) Run(command string, args ...string) error {
	ce.Command = command
	ce.Args = args
	//util.Logger.Info("args", zap.Any("args", args))
	cmd := exec.Command(ce.Command, ce.Args...)
	cmd.Stdout = &ce.Stdout
	cmd.Stderr = &ce.Stderr
	fmt.Println("Command:", cmd.String())
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
		return err
	}

	return nil
}

// Output 获取标准输出
func (ce *CommandExecutor) Output() string {
	return ce.Stdout.String()
}

// StderrOutput 获取标准错误输出
func (ce *CommandExecutor) StderrOutput() string {
	return ce.Stderr.String()
}

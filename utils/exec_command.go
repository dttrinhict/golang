package utils

import (
	"os"
	"os/exec"
)

const ShellToUse = "bash"
/* Exec bash command
https://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go
*/
func ExecCommand(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	// Open standard input - allow input from keyboard
	cmd.Stdin = os.Stdin
	// Open standard output - show command's result
	cmd.Stdout = os.Stdout
	// Open standard error - show command's error
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
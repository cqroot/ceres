package utils

import (
	"bufio"
	"fmt"
	"os/exec"
)

func ExecCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	return err
}

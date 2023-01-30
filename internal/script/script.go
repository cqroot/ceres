package script

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Run(scriptPath string, workDir string) error {
	fmt.Println("Run script:")
	fmt.Println()
	fmt.Println("    dir:", workDir)
	fmt.Println("    cmd:", []string{"/bin/bash", scriptPath})
	fmt.Println()

	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Dir = workDir

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	return err
}

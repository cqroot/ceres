package script

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/cqroot/ceres/internal/utils"
)

func Run(scriptPath string, workDir string) error {
	fmt.Println(utils.ColorString("Run script:"))
	fmt.Println()
	fmt.Println(utils.ColorString("    dir:"), workDir)
	fmt.Println(utils.ColorString("    cmd:"), []string{"/bin/bash", scriptPath})
	fmt.Println()

	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Dir = workDir

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

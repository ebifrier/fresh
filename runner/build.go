package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building...")
	buildLog("options: -gcflags=%s -ldflags=%s -o=%s test_%s",
		gcflags(), ldflags(), buildPath(), buildRoot())

	cmd := exec.Command(
		"go", "build", "-gcflags", gcflags(), "-ldflags", ldflags(),
		"-o", buildPath(), buildRoot())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}

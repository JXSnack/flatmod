package gradle

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/JXSnack/flatmod/idea"
)

// ExecGradlew executes a gradle command on the Workspace.
func (w *Workspace) ExecGradlew(gradleArgs []string) error {
	wrapper := "gradlew"
	if runtime.GOOS == "windows" {
		wrapper = "gradlew.bat"
	}

	wrapperPath, _ := filepath.Abs(filepath.Join(w.Dir(), wrapper))

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", append([]string{"/C", wrapperPath}, gradleArgs...)...)
	} else {
		cmd = exec.Command(wrapperPath, gradleArgs...)
	}

	jdk, err := idea.GetJDK()
	if err != nil {
		return err
	}

	cmd.Env = append(os.Environ(), "JAVA_HOME="+jdk)
	cmd.Dir = w.Dir()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

package modes

import (
	"os"
	"os/exec"
)


func InitializeGoModule(projectPath string ) error {
	exec.Command("cd",  projectPath)

	projectPath = "github.com/" + projectPath

	cmd := exec.Command("go", "mof", "init", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
}
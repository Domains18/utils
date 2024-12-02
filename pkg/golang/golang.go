package golang

import (
	"fmt"
	"os"
	"os/exec"
)


func InitializeGoModule(projectPath string ) error {
	exec.Command("cd",  projectPath)

	projectPath = "github.com/" + projectPath

	cmd := exec.Command("go", "mod", "init", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running go mod init")
	}

	fmt.Printf("Project initialized")
	exec.Command(("cd .."))

	return nil
}


func InstallLibraries(libraries  string) error {
	cmd := exec.Command("go", "get", "-u", libraries)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running go get %v", err)
	}
	return nil
}
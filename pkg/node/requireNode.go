package pkg

import (
	"fmt"
	"os"
	"os/exec"
)



type express struct{
	projectName string
	inntializeGit bool
}


func CreateExpressProject (projectName string, inntializeGit bool) express{
	project := express{ projectName: projectName, inntializeGit: inntializeGit}

	err := os.Mkdir(projectName, os.ModePerm)
	if err != nil {
		fmt.Printf("error initializing the project : %v\n", err)
		os.Exit(1)
	}

	err = os.Chdir(projectName)
	if err != nil {
		fmt.Printf("error changing directory to project directory %v\n", err)
	}
	
	cmd := exec.Command("npm", " init", "-y")
	err = cmd.Run()
	if err != nil{
		fmt.Printf("error initializing ")
	}

	return project
}


func RequireNode(){

}
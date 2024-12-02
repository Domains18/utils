package projectstructure

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Domains18/utils/pkg/golang"
	"github.com/cheggaaa/pb/v3"
)



var directories = []string {
	"cmd",
	"internal/api/handlers",
	"internal/app/config",
	"internal/app/database/postgres",
	"internal/app/database/mysql",
	"internal/app/middleware",
	"api/v1/routes",
	"scripts",
	"web",
}

var files = []string{
	"Dockerfile",
		"go.mod",
		"go.sum",
		"README.md",
		"cmd/main.go",
		"internal/api/api.go",
		"internal/app/app.go",
		"internal/app/database/postgres/postgres.go",
		"internal/app/database/mysql/mysql.go",
		"api/v1/routes/routes.go",
}


var libraries = []string{
	"github.com/gorilla/mux",
	"github.com/joho/godotenv",
	"github.com/jinzhu/gorm",
	"github.com/jinzhu/gorm/dialects/postgres",
	"github.com/jinzhu/gorm/dialects/mysql",
	"github.com/dgrijalva/jwt-go",
}


func CreateProjectStructure(projectName string, useDocker bool){

	if _, err := os.Stat(projectName); !os.IsNotExist(err){
		log.Fatalf("projectName with %s already exists, delete or choose a different name", projectName)
	}
	if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}

	if err := os.Chdir(projectName); err != nil {
		log.Fatalln("failed to change directory")
	}

	//create go module
	if err :=  golang.InitializeGoModule(projectName); err != nil {
		log.Fatalf("failed to initialize the go modules %v", err)
	}

	dirsProgressBar := pb.StartNew((len(directories)))
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatalf("failed to create directories %v", err)
		}
		dirsProgressBar.Increment()
	}
	dirsProgressBar.Finish()
	

	filesProgress := pb.StartNew(len(files))
	for _, file := range files {
		if _, err := os.Create(file); err != nil {
			log.Fatalf("failed to create directories %v", err)
		}
		filesProgress.Increment()
	}
	filesProgress.Finish()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory %v", err)
	}

	parentDir := filepath.Dir(cwd)
	projectDir := filepath.Join(parentDir, projectName)
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	fmt.Print("\x1b[34mHere are the libraries you have to install: \x1b[0m\n")
	
	for i, lib := range libraries {
		fmt.Printf("\x1b[32m%d. %s\x1b[0m\n", i+1, lib)
	}
	if i := len(libraries); i > 0 {
		fmt.Printf("\x1b[34mRun the following command to install the libraries: \x1b[0m\n")
		fmt.Printf("\x1b[32mgo get %s\x1b[0m\n", libraries)
	}
	
}
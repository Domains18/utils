package projectstructure

import (
	"log"
	"os"
)




func createProjectStructure(projectName string, useDocker bool){

	if _, err := os.Stat(projectName); !os.IsNotExist(err){
		log.Fatalln("projectName with %s already exists, delete or choose a different name", projectName)
	}
	if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
		log.Fatalln("failed to create directory: %v", err)
	}

	if err := os.Chdir(projectName); err != nil {
		log.Fatalln("failed to change directory")
	}

	//create go module
}
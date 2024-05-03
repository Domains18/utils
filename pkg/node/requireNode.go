package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

type express struct {
	projectName   string
	inntializeGit bool
}

var packages = []string{
	"express",
	"body-parser",
	"cors",
	"dotenv",
	"nodemon",
	"axios",
	"bcrypt",
	"cookie-parser",
	"jsonwebtoken",
	"http",
	"prisma",
}

var devPackages = []string{
	"eslint",
	"prettier",
	"@typescript-eslint/parser",
	"@typescript-eslint/eslint-plugin",
	"typescript",
	"ts-node",
	"ts-node-dev",
	"@types/express",
	"@types/bcrypt",
	"@types/cookie-parser",
	"@types/jsonwebtoken",
	"@types/node",
}

var scripts = map[string]string{
	"dev": "nodemon src/index.js",
	"start": "node src/index.js",
	"build": "tsc",
	"lint": "eslint . --ext .ts",
	"format": "prettier --write .",
}

var gitignore = []string{
	"node_modules",
	".env",
	".DS_Store",
	".vscode",
	".idea",
	"dist",
	"build",
}

var tsconfig = map[string]interface{}{
	"compilerOptions": map[string]interface{}{
		"target": "es6",
		"module": "commonjs",
		"outDir": "./dist",
		"rootDir": "./src",
		"strict": true,
		"esModuleInterop": true,
		"skipLibCheck": true,
		"forceConsistentCasingInFileNames": true,
		"resolveJsonModule": true,
		"baseUrl": "./src",
		"paths": map[string]interface{}{
			"*": ["node_modules/*", "src/types/*"],
		},
	},
	"include": ["src/**/*.ts"],
	"exclude": ["node_modules"],
}

var eslint = map[string]interface{}{
	"root": true,
	"parser": "@typescript-eslint/parser",
	"plugins": ["@typescript-eslint"],
	"extends": [
		"eslint:recommended",
		"plugin:@typescript-eslint/recommended",
		"plugin:@typescript-eslint/recommended-requiring-type-checking",
	],
	"parserOptions": {
		"project": "./tsconfig.json",
	},
	"rules": {
		"@typescript-eslint/no-explicit-any": "off",
		"@typescript-eslint/explicit-module-boundary-types": "off",
		"@typescript-eslint/no-unsafe-assignment": "off",
		"@typescript-eslint/no-unsafe-member-access": "off",
		"@typescript-eslint/no-unsafe-call": "off",
		"@typescript-eslint/no-unsafe-return": "off",
		"@typescript-eslint/restrict-template-expressions": "off",
		"@typescript-eslint/no-unsafe-return": "off",
		"@typescript-eslint/no-unsafe-assignment": "off",
		"@typescript-eslint/no-unsafe-member-access": "off",
		"@typescript-eslint/no-unsafe-call": "off",
		"@typescript-eslint/no-floating-promises": "off",
		"@typescript-eslint/no-misused-promises": "off",
		"@typescript-eslint/no-unsafe-return": "off",
		"@typescript-eslint/no-unsafe-assignment": "off",
		"@typescript-eslint/no-unsafe-member-access": "off",
		"@typescript-eslint/no-unsafe-call": "off",
		"@typescript-eslint/no-floating-promises": "off",
	}
}

var projectStructure = []string{
	"src",
	"src/controllers",
	"src/middlewares",
	"src/models",
	"src/routes",
	"src/services",
	"src/types",
	"src/utils",
	"src/index.ts",
}

func CreateExpressProject(projectName string, inntializeGit bool) express {
	project := express{projectName: projectName, inntializeGit: inntializeGit}

	err := os.Mkdir(projectName, os.ModePerm)
	if err != nil {
		fmt.Printf("error initializing the project : %v\n", err)
		os.Exit(1)
	}

	err = os.Chdir(projectName)
	if err != nil {
		fmt.Printf("error changing directory to project directory %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("npm", "init", "-y")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error initializing the project with npm: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("installing packages...................")
	cmd = exec.Command("npm", append([]string{"install", "-s"}, packages...)...)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error installing packages: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("installing dev packages...................")
	cmd = exec.Command("npm", append([]string{"install", "-D"}, devPackages...)...)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error installing dev packages: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("adding gitingore file...................")
	file, err := os.Create(".gitignore")
	if err != nil {
		fmt.Printf("error creating gitignore file: %v\n", err)
		os.Exit(1)
	}
	for _, line := range gitignore {
		file.WriteString(line + "\n")
	}
	file.Close()
	fmt.Printf("adding scripts to package.json...................")
	for key, value := range scripts {
		cmd = exec.Command("npm", "set-script", key, value)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("error adding scripts to package.json: %v\n", err)
			os.Exit(1)
		}
	}

}

func RequireNode() {

}

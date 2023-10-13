package main

import (
	"fmt"
	"os"

	"os/exec"
)

func file_exists(file string) (bool, error) {

	_, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading %s file: %v\n", file, err)
		return false, err
	}
	return true, nil

}

func main() {

	// modFile, err := modfile.Parse(goModPath, data, nil)
	// if err != nil {
	// 	fmt.Printf("Error parsing go.mod file: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Module name: %s\n", modFile.Module.Mod.Path)

	// fmt.Println("Require directives:")
	// for _, require := range modFile.Require {
	// 	fmt.Printf("%s %s\n", require.Mod.Path, require.Mod.Version)
	// }

	// fmt.Println("Replace directives:")
	// for _, replace := range modFile.Replace {
	// 	fmt.Printf("Old: %s New: %s\n", replace.Old.Path, replace.New.Path)
	// }

	// cmd := exec.Command("go", "list", "-m", "-u", "all")

	// cmdDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("Error getting current working directory: %v\n", err)
	// 	return
	// }

	// cmd.Dir = cmdDir

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("Error running go list: %v\n", err)
	// 	return
	// }

	// fmt.Printf("go list output:\n%s\n", string(output))

	_, err := file_exists("/workspace/go.sum")
	if err != nil {
		return
	}

	cmdGet := exec.Command("go", "get", "-u", "./...")

	cmdGet.Dir = os.Args[1]

	outputGet, err := cmdGet.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running go get: %v\n", err)
		return
	}

	fmt.Printf("go get -u ./...\n%s\n", string(outputGet))

	cmdMod := exec.Command("go", "mod", "tidy")

	cmdMod.Dir = os.Args[1]

	outputMod, err := cmdMod.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running go mod tidy: %v\n", err)
		return
	}

	fmt.Printf("go mod tidy\n%s\n", string(outputMod))
}

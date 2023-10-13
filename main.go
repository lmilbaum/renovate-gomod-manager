package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	goModPath := "go.mod"

	data, err := os.ReadFile(goModPath)
	if err != nil {
		fmt.Printf("Error reading go.mod file: %v\n", err)
		return
	}

	modFile, err := modfile.Parse(goModPath, data, nil)
	if err != nil {
		fmt.Printf("Error parsing go.mod file: %v\n", err)
		return
	}

	fmt.Printf("Module name: %s\n", modFile.Module.Mod.Path)

	fmt.Println("Require directives:")
	for _, require := range modFile.Require {
		fmt.Printf("%s %s\n", require.Mod.Path, require.Mod.Version)
	}

	fmt.Println("Replace directives:")
	for _, replace := range modFile.Replace {
		fmt.Printf("Old: %s New: %s\n", replace.Old.Path, replace.New.Path)
	}
}

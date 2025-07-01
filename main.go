package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

const VIDEOS_PATH = "C:/Users/delse/Videos/"

func createSubFolders(basePath string, subFolders map[string][]string) {
	for folder, nested := range subFolders {
		subFolderPath := basePath + "/" + folder

		err := os.MkdirAll(subFolderPath, os.ModePerm)

		if err != nil {
			color.Red("Error creating subfolder %s: %v", folder, err)
		} else {
			color.Green("Created subfolder: %s", subFolderPath)
		}

		// Create nested subfolders if any
		for _, n := range nested {
			nestedPath := subFolderPath + "/" + n

			err := os.MkdirAll(nestedPath, os.ModePerm)

			if err != nil {
				color.Red("Error creating nested subfolder %s: %v", n, err)
			} else {
				color.Green("Created nested subfolder: %s", nestedPath)
			}
		}
	}
}

func createFolderStructure(projectName string) {
	path := VIDEOS_PATH + projectName

	// Check if the folder already exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		color.Red("Folder already exists!")
		return
	}

	// Create the folder structure
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating folder structure:", err)
		return
	}

	fmt.Println("Folder created at:", path)

	// Create 4 subdirectories
	subFolders := map[string][]string{
		"00 Project Info":  {},
		"01 Assets":        {"01 FOOTAGE", "02 ASSETS", "03 SFX", "04 MUSIC", "05 VFX", "06 GRAPHICS", "07 IMAGES"},
		"02 Project Files": {},
		"03 Exports":       {"Drafts", "Final"},
	}
	createSubFolders(path, subFolders)
}

func appendDateToProjectName(projectName string) string {
	currentDate := time.Now().Format("2 January 2006")
	return projectName + " - " + currentDate
}

func main() {
	projectName := flag.String("name", "", "Project name for the folder structure")
	help := flag.Bool("help", false, "Display help information")

	flag.Parse()

	if *projectName == "" {
		fmt.Print("Please provide a project name using -name flag.\n")
		return
	}

	if *help {
		color.Green("Welcome to the Video Template Creator!")
		color.Yellow("This tool will help you create a folder structure for your video projects.")
	}

	createFolderStructure(appendDateToProjectName(*projectName))
}

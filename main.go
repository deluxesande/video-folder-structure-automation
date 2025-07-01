package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const VIDEOS_PATH = "C:/Users/delse/Videos/"
const TEMPLATE_PATH = VIDEOS_PATH + "Video Template/"

func createFolderStructure(projectName string) {
	var path = projectName
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

	fmt.Println("Folder structure created at:", path)
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

	createFolderStructure(*projectName)
}

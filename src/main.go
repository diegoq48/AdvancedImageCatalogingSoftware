package AdvancedImageCatalogingSoftware

import (
	"fmt"
	"os"
)

func main() {
	// safe os.args into a variable
	args := os.Args[1:]
	// here is a list of options if we see the --catalog option we will create a catalog so the next argument after catalog will be a filepath to a directory
	// if we see a --find options the next option will be a stored catalog followed by a csv of all of the tags to look for
	// if we see a --help option we will print out the help menu
	// if we see a --version option we will print out the version of the program
	// if we see a --list option we will print out all of the stored catalogs
	// if we see a --list-tags option we will print out all of the tags in a stored catalog
	// if we see a --list-files option we will print out all of the files in a stored catalog
	// if we see a --correct option we will correct the tags in a stored catalog by allowing the user to manualy edit the tags assosiated with a file so it should be followed by a filepath
	// if we see a --remove option we will remove a stored catalog so it should be followed by a filepath

	// if there are no arguments print out the help menu
	if len(args) == 0 {
		printHelpMenu()
		return
	}
	switch args[0] {
	case "--catalog":
		if len(args) < 2 {
			fmt.Println("Please specify a directory to catalog")
			return
		}
		if !isDirectory(args[1]) {
			fmt.Println("Please specify a directory to catalog")
			return
		}
		if !acessingFiles(args[1]) {
			fmt.Println("Error tagging files")
			return
		}
		acessingFiles(args[1])
		fmt.Println("Catalog created")
	case "--help":
		printHelpMenu()
	case "--version":
		fmt.Println("Version 1.0")
	case "--list":
		fmt.Println("List of catalogs")
		// list all the files in /opt/advancedImageCatalogingSoftware/catalogs
		for file in os.Open("/opt/advancedImageCatalogingSoftware/catalogs") {
			fmt.Println(file);
		}
	case "--list-tags":
		if len(args) < 2 {
			fmt.Println("Please specify a catalog")
			return
		}
		//check if the catalog exists in /opt/advancedImageCatalogingSoftware/catalogs
		if !fileExists("/opt/advancedImageCatalogingSoftware/catalogs/" + args[1]) {
			fmt.Println("Please specify a catalog present in /opt/advancedImageCatalogingSoftware/catalogs")
			return;
		}
		// list all the tags in the catalog
		
	
	}

}

func printHelpMenu() {
	fmt.Println(
		"--catalog <directory> \t creates a catalog of all the files in the directory specified in the filepath and stores it for future use",
		"--find <filepath || catalog name> <tag1,tag2,tag3...> \t finds all of the files in the catalog specified in the filepath that have all of the tags specified",
		"--help \t prints out the help menu",
		"--version \t prints out the version of the program",
		"--list \t prints out all of the stored catalogs",
		"--list-tags <catalog name> \t prints out all of the tags in the catalog specified in the filepath",
		"--list-files <catalog name> \t prints out all of the files in the catalog specified in the filepath",
		"--correct <catalog name> <realfilepath> \t corrects the tags in the catalog specified in the filepath",
		"--remove <catalog name> \t removes the catalog specified in the arguments",
	)
}

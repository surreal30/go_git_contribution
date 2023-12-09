package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
	"bufio"
)

func scan(path string) {
	print("scan")
	fmt.Printf("Found folders: \n \n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("Successfully added \n")
}

func scanGitFolders(folders []string, folder string) []string {
	folder = string.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = path + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue			
			}

			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}

	}

	return folders 
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0) folder)
}

func getDotFilePath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

type User struct {
	Uid string 
	Gid string 
	Username string
	Name string
	HomeDir string
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSlicesToFile(repos, filePath)
}

func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
}

func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for the git repositories")
	flag.StringVar(&email, "email", "example@email.com", "the email to scan")

	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
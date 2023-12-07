package main

import (
	"flag"
)

func scan(path string) {
	print("scan")
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
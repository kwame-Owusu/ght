package cmd

import (
	"fmt"
	"strings"
)


type RepoString struct {
	userName string
	projectName string
}


func parseRepoURL(repo string) RepoString{
	if len(repo) == 0 {
		fmt.Println("Error: empty repo URL string")
		return RepoString{}
	}

	input := strings.Trim(repo, "/")
	parts := strings.Split(input, "/")

	if len(parts) < 2 {
		fmt.Println("Error: Invalid format. Use 'user/repo' (e.g., ght get golang/go)")
		return RepoString{}
	}

	return	RepoString{userName: parts[0], projectName: parts[1]}
}
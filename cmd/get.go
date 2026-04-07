package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Clone trending GitHub repository",
	Run:   cloneTrendingRepo,
}

func cloneTrendingRepo(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a repository (e.g., ght get user/repo)")
		return
	}

	input := strings.Trim(args[0], "/")
	parts := strings.Split(input, "/")

	if len(parts) < 2 {
		fmt.Println("Error: Invalid format. Use 'user/repo' (e.g., ght get golang/go)")
		return
	}

	userName := parts[0]
	projectName := parts[1]
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", userName, projectName)

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error finding home directory: %v\n", err)
		return
	}

	fmt.Printf("Cloning %s/%s...\n", userName, projectName)

	gitCloneCmd := exec.Command("git", "clone", repoURL)
	gitCloneCmd.Dir = home

	output, err := gitCloneCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Clone failed:\n%s", string(output))
		return
	}

	fmt.Printf(" Successfully cloned to %s/%s\n", home, projectName)
}

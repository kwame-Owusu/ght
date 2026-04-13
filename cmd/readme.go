package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)


var getReadme = &cobra.Command{
	Use:   "readme",
	Short: "Get the readme of repo on terminal",
	Run:   repoReadme,
}

type ReadmeResponse struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

func repoReadme(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a repository (e.g., ght get user/repo)")
		return
	}

	repo := parseRepoURL(args[0])
	readme, err := fetchReadme(repo.userName, repo.projectName)
	if err != nil{
		fmt.Println("Error: error fetching repo README")
		return
	}

	out, err := glamour.Render(readme, "dark") 
	if err != nil{
		fmt.Println("Error: error rendering markdown through glamour")
		return
	}

	execCmd := exec.Command("less", "-R") // -R = allow ANSI colors
	execCmd.Stdin = strings.NewReader(out)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	execCmd.Run()
}

func fetchReadme(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/readme", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "ght")
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API %s for %s/%s: %s", resp.Status, owner, repo, body)
	}

	var data ReadmeResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("decoding JSON: %w", err)
	}

	if data.Encoding != "base64" {
		return "", fmt.Errorf("unexpected encoding %q (expected base64)", data.Encoding)
	}

	// GitHub base64-encodes content with line breaks every 60 chars.
	// StdEncoding rejects embedded newlines, so strip them first.
	cleaned := strings.ReplaceAll(data.Content, "\n", "")
	decoded, err := base64.StdEncoding.DecodeString(cleaned)
	if err != nil {
		return "", fmt.Errorf("decoding base64: %w", err)
	}

	return string(decoded), nil
}
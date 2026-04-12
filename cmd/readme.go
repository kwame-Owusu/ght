package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	fmt.Print(out)
	fmt.Print("\033[H")
}

func fetchReadme(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/readme", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// GitHub recommends setting a User-Agent
	req.Header.Set("User-Agent", "ght")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed: %s\n%s", resp.Status, body)
	}

	var data ReadmeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	// Decode Base64 content
	decoded, err := base64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

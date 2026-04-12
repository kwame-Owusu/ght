package cmd

import (
	"github.com/spf13/cobra"
)


var getReadme = &cobra.Command{
	Use:   "readme",
	Short: "Get the readme of repo on terminal",
	Run:   getRepoReadme,
}


func getRepoReadme(cmd *cobra.Command, args []string) {
	// url := fmt.Sprint("https://githubusercontent.com[User]/[Repo]/[Branch]/[FilePath")
}
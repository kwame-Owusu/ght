package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/kwame-Owusu/ght/internal"
	"github.com/spf13/cobra"
)

var dayFlag bool

var rootCmd = &cobra.Command{
	Use:   "ght",
	Short: "Get trending github repos",
	Run:   getTrending,
}

func getTrending(cmd *cobra.Command, args []string) {
	internal.MakeDayRequest()
}

func init() {
	rootCmd.Flags().BoolVarP(&dayFlag, "day", "d", false, "Get trending today")
}

func Execute() {
	ctx := context.Background()
	if err := fang.Execute(ctx, rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

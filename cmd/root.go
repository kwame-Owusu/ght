package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/kwame-Owusu/ght/internal"
	"github.com/spf13/cobra"
)

var dayFlag, weekFlag, monthFlag bool

var rootCmd = &cobra.Command{
	Use:   "ght",
	Short: "Get trending github repos",
	Run:   getTrending,
}

func getTrending(cmd *cobra.Command, args []string) {
	switch {
	case dayFlag:
		internal.MakeDayRequest()
	case weekFlag:
		internal.MakeWeekRequest()
	case monthFlag:
		internal.MakeMonthRequest()
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&dayFlag, "day", "d", false, "Get trending repos today")
	rootCmd.Flags().BoolVarP(&weekFlag, "week", "w", false, "Get trending repos for the week")
	rootCmd.Flags().BoolVarP(&monthFlag, "month", "m", false, "Get trending repos for this month")
}

func Execute() {
	ctx := context.Background()
	if err := fang.Execute(ctx, rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

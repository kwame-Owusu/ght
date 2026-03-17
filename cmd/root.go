package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/kwame-Owusu/ght/internal"
	"github.com/spf13/cobra"
)

var dayFlag, weekFlag, monthFlag, yearFlag bool
var limitFlag int

var rootCmd = &cobra.Command{
	Use:   "ght",
	Short: "Get trending github repos",
	Run:   getTrending,
}

func getTrending(cmd *cobra.Command, args []string) {
	params := internal.TrendingParams{
		DayFlag:   dayFlag,
		WeekFlag:  weekFlag,
		MonthFlag: monthFlag,
		YearFlag:  yearFlag,
		Limit:     limitFlag,
	}

	internal.MakeTrendingRequest(params)
}

func init() {
	rootCmd.Flags().BoolVarP(&dayFlag, "day", "d", false, "Get trending repos today")
	rootCmd.Flags().BoolVarP(&weekFlag, "week", "w", false, "Get trending repos for the week")
	rootCmd.Flags().BoolVarP(&monthFlag, "month", "m", false, "Get trending repos for this month")
	rootCmd.Flags().BoolVarP(&yearFlag, "year", "y", false, "Get trending repos for past year")
	rootCmd.Flags().IntVarP(&limitFlag, "limit", "l", 5, "Limit the response of github repos")
	rootCmd.MarkFlagsMutuallyExclusive("day", "week", "month")
}

func Execute() {
	ctx := context.Background()
	if err := fang.Execute(ctx, rootCmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

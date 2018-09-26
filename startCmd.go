package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new timetracking interval",
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := loadTimeSheet(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := ts.Start(time.Now()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := ts.Save(path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		ts.Print(os.Stdout, time.Duration(roundToMinutes)*time.Minute)
	},
}
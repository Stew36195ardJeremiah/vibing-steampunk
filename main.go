package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// rootCmd is the base command for vibing-steampunk
var rootCmd = &cobra.Command{
	Use:   "vibing-steampunk",
	Short: "A steampunk-themed vibing tool",
	Long: `vibing-steampunk is a tool for generating steampunk-themed
content and visualizations. Embrace the gears, the steam, and the vibe.

Personal fork: tweaked for my own experimentation and learning.
See also: https://github.com/oisee/vibing-steampunk (upstream)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// versionCmd prints the current version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("vibing-steampunk %s (commit: %s, built: %s)\n", version, commit, date)
		fmt.Println("(personal fork — upstream: github.com/oisee/vibing-steampunk)")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

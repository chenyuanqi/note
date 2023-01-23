package controllers

import (
	"mlj/pkg/console"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	console.Success("hello")
}

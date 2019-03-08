package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "email",
	Short: "Send Emails from your command line",
}

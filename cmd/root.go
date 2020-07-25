package cmd

import "github.com/spf13/cobra"

// RootCmd is a root Cobra command that gets called
// from the main func. All other sub-commands should
// be registered here.
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "passgen",
		Short: "XKCD inspired CLI password generator",
		Long: `A command line tool for generating
passwords inspired by an XKCD commic:
https://xkcd.com/936/`,
	}
	cmd.AddCommand(
		generateCmd(),
	)
	return cmd
}

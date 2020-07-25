package cmd

import (
	"github.com/spf13/cobra"

	"github.com/antolis/passgen/internal/app"
)

func generateCmd() *cobra.Command {
	a := app.New()
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate password",
		Long: `Generate a password by appending several
randomly chosen common english words.`,
		Example: `Generate a password without much fuss:
$ passgen generate

Generate a password with at least 24 characters, both
lower and upper case letters and special characters:
$ passgen generate -scm 24

Or store the same config into a file and reuse it for
all generated passwords:
$ echo 'capital-letters = true
special-chars = true
min-length = 24' > ~/.config/passgen.toml
$ passgen generate

Supported file formats:
json, yaml, toml, hcl & java props.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initParams(cmd, a.Params)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Generate()
		},
	}
	cmd.Flags().IntVarP(&a.Params.MinLength, "min-length", "m", 16, "Specify minimum password length, must not be less than 8")
	cmd.Flags().BoolVarP(&a.Params.SpecialChars, "special-chars", "s", false, "Request non-alphanumeric characters to be included in the password")
	cmd.Flags().BoolVarP(&a.Params.CapitalLetters, "capital-letters", "c", false, "Require password to contain both lowercase and uppercase letters")
	return cmd
}

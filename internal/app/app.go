/*
Package app contains all the business logic for
generating passwords. It has no dependency on either
viper or cobra libraries, and can be unit tested.

Output writter for the App is configurable so that
tests can easily capture the output and perform
assertions on it. Params are initialized in the param
package, so that app can be clear of viper dependency.
*/
package app

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

// App implements the core logic of generating passwords.
type App struct {
	// Parameters passed to the CLI by the user.
	Params *Params
	// Out is used to write the output. Defaults to
	// standard out, but can be overridden in tests
	// to make assertions on the application's output.
	Out io.Writer
	// Source of randomness for the application.
	// Tests can use a mocked random source in order
	// to provide a deterministic testing behavior.
	Random io.Reader
}

// Params struct holds all the user customizable
// parameters for the application. Using a single struct
// for all CLI commands ensures that all flags are distinct
// and that they can be provided either dynamically on a
// command line, or statically in a config file that's
// reused between command runs.
type Params struct {
	MinLength      int  `mapstructure:"min-length"`
	SpecialChars   bool `mapstructure:"special-chars"`
	CapitalLetters bool `mapstructure:"capital-letters"`
}

// New instantiates an App with default parameters,
// including standard output and cryptographically
// secure random source.
func New() *App {
	return &App{
		Params: &Params{},
		Out:    os.Stdout,
		Random: rand.Reader,
	}
}

// Generate the password based on the current params
// in the app instance.
func (a *App) Generate() error {
	if err := a.validateParams(); err != nil {
		return err
	}
	var pass string
	var passSize int
	var words [2048]string = words()
	for passSize < a.Params.MinLength {
		i, err := rand.Int(a.Random, big.NewInt(2047))
		if err != nil {
			return err
		}
		next := words[i.Int64()]
		if a.Params.CapitalLetters {
			next = strings.Title(next)
		}
		if passSize > 0 && a.Params.SpecialChars {
			next = " " + next
		}
		pass += next
		passSize = size(pass)
	}
	if a.Params.SpecialChars {
		pass += "!"
	}
	a.Out.Write([]byte(pass))
	return nil
}

func (a *App) validateParams() error {
	params := a.Params
	if params.MinLength < 8 {
		return fmt.Errorf("min-length (%d) must not be less than 8", params.MinLength)
	}
	return nil
}

func size(word string) int {
	return len([]rune(word))
}

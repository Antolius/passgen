package app_test

import (
	"bytes"
	"testing"

	"github.com/antolis/passgen/internal/app"
)

func TestParamValidation(t *testing.T) {
	app := app.New()
	app.Params.MinLength = 7
	err := app.Generate()
	if err == nil {
		t.Fatalf("expected app to return error but it didn't")
	}
}

func TestPasswordGeneration(t *testing.T) {
	tests := map[string]struct {
		input app.Params
		want  string
	}{
		"simple":          {app.Params{8, false, false}, "abilityable"},
		"long":            {app.Params{16, false, false}, "abilityableabout"},
		"special chars":   {app.Params{8, true, false}, "ability able!"},
		"capital letters": {app.Params{8, false, true}, "AbilityAble"},
		"all features":    {app.Params{16, true, true}, "Ability Able About!"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			// rand will produce a random int sequence 1, 2, 3
			rand := bytes.NewReader([]byte{0, 0, 0, 1, 0, 2})
			app := app.App{
				Params: &test.input,
				Out:    buf,
				Random: rand,
			}
			err := app.Generate()
			if err != nil {
				t.Fatalf("expected no errors, but got %#v", err)
			}
			out := buf.String()
			if out != test.want {
				t.Fatalf("expected %#v, but got %#v", test.want, out)
			}
		})
	}
}

func TestRandomError(t *testing.T) {
	app := app.App{
		Params: &app.Params{MinLength: 8},
		Out:    new(bytes.Buffer),
		Random: bytes.NewReader([]byte{0}),
	}
	err := app.Generate()
	if err == nil {
		t.Fatalf("expected app to return error but it didn't")
	}
}

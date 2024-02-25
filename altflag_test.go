package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

type altFlagTest interface {
	clargs() []string
	expectedErrorStringContaining() *string
	flagSetName() *string
	setupFlagSet(f altflag.FlagSet)
	verify(t *testing.T, f altflag.FlagSet)
}

func TestParse(t *testing.T) {
	for name, cfg := range map[string]altFlagTest{
		"base": newAltFlagTestNoop(),
		"simple string var": newAltFlagTestSimpleStringVar(
			"someStringVariable",
			"s",
			"Some usage string",
			[]string{
				"--someStringVariable",
				"some-string-value",
			},
			ptr("some-string-value"),
			nil,
		),
		"simple string var no match": newAltFlagTestSimpleStringVar(
			"foo",
			"f",
			"Some usage string",
			[]string{
				"--bar",
				"some-value",
			},
			nil,
			ptr("argument --bar didn't match any known flags"),
		),
		"string var multimatch": newAltFlagTestMultiStringVar(
			[]string{"foobaz", "foobar"},
			[]string{"a", "b"},
			[]string{"Some usage string", "Some other usage string"},
			[]string{
				"--foo",
				"some-value",
			},
			nil,
			ptr("argument --foo matched multiple flags: --foobar, --foobaz"),
		),
	} {
		t.Run(name, newAltFlagTestFunc(cfg))
	}
}

func newAltFlagTestFunc(cfg altFlagTest) func(t *testing.T) {
	return func(t *testing.T) {
		flagSetName := cfg.flagSetName()
		if flagSetName == nil {
			flagSetName = ptr("some-flagset")
		}
		flagSet := altflag.NewFlagSet(*flagSetName)
		cfg.setupFlagSet(flagSet)
		err := flagSet.Parse(cfg.clargs())
		expectedErrorStringContaining := cfg.expectedErrorStringContaining()
		if expectedErrorStringContaining != nil {
			if assert.Error(t, err, "Parse should return an error") {
				assert.ErrorContains(t, err, *expectedErrorStringContaining)
			}
		} else {
			assert.Nil(t, err, "Parse should return nil")
			cfg.verify(t, flagSet)
		}
	}
}

func ptr[T any](value T) *T {
	return &value
}

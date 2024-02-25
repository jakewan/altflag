package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

type altFlagTest interface {
	clargs() []string
	expectedSetupErrorStringContaining() *string
	expectedParseErrorStringContaining() *string
	flagSetName() *string
	setupFlagSet(f altflag.FlagSet) error
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
			nil,
			ptr("argument --foo matched multiple flags: --foobar, --foobaz"),
		),
		"duplicate string vars": newAltFlagTestMultiStringVar(
			[]string{"foo", "foo"},
			[]string{"a", "b"},
			[]string{"Some usage string", "Some other usage string"},
			nil,
			nil,
			ptr("argument --foo is already configured"),
			nil,
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
		err := cfg.setupFlagSet(flagSet)
		expectedParseErrorStringContaining := cfg.expectedSetupErrorStringContaining()
		if expectedParseErrorStringContaining != nil {
			if assert.Error(t, err, "setupFlagSet should return an error") {
				assert.ErrorContains(t, err, *expectedParseErrorStringContaining)
			}
			return
		} else {
			assert.Nil(t, err, "setupFlagSet should return nil")
		}

		err = flagSet.Parse(cfg.clargs())
		expectedParseErrorStringContaining = cfg.expectedParseErrorStringContaining()
		if expectedParseErrorStringContaining != nil {
			if assert.Error(t, err, "Parse should return an error") {
				assert.ErrorContains(t, err, *expectedParseErrorStringContaining)
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

package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

type altFlagTest interface {
	clargs() []string
	expectedError() error
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
		cfg.setupFlagSet(flagSet)
		err := flagSet.Parse(cfg.clargs())
		expectedError := cfg.expectedError()
		if expectedError != nil {
			if assert.Error(t, err, "Parse should return an error") {
				assert.Equal(t, cfg.expectedError, err)
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

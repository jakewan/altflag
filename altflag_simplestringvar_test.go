package altflag_test

import (
	"testing"

	"github.com/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

func newAltFlagTestSimpleStringVar(
	displayName string,
	shortFlag string,
	usage string,
	clargs []string,
	expectedValue *string,
	setupErrorStringContaining *string,
	parseErrorStringContaining *string) altFlagTest {
	return &altFlagTestSimpleStringVar{
		displayName:                displayName,
		shortFlag:                  shortFlag,
		usage:                      usage,
		testClargs:                 clargs,
		expectedValue:              expectedValue,
		setupErrorStringContaining: setupErrorStringContaining,
		parseErrorStringContaining: parseErrorStringContaining,
	}
}

type altFlagTestSimpleStringVar struct {
	displayName                string
	shortFlag                  string
	usage                      string
	myVar                      string
	testClargs                 []string
	expectedValue              *string
	setupErrorStringContaining *string
	parseErrorStringContaining *string
}

// expectedSetupErrorStringContaining implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) expectedSetupErrorStringContaining() *string {
	return aft.setupErrorStringContaining
}

// expectedParseErrorStringContaining implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) expectedParseErrorStringContaining() *string {
	return aft.parseErrorStringContaining
}

// clargs implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) clargs() []string {
	return aft.testClargs
}

// flagSetName implements altFlagTest.
func (*altFlagTestSimpleStringVar) flagSetName() *string { return nil }

// setupFlagSet implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) setupFlagSet(f altflag.FlagSet) error {
	_, err := f.StringVar(&aft.myVar, aft.displayName, aft.shortFlag, aft.usage)
	return err
}

// verify implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) verify(t *testing.T, f altflag.FlagSet) {
	assert.Equal(t, aft.myVar, *aft.expectedValue)
}

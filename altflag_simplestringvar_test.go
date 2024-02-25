package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

func newAltFlagTestSimpleStringVar(
	displayName string,
	shortFlag string,
	usage string,
	clargs []string,
	expectedValue *string,
	errorStringContaining *string) altFlagTest {
	return &altFlagTestSimpleStringVar{
		displayName:           displayName,
		shortFlag:             shortFlag,
		usage:                 usage,
		testClargs:            clargs,
		expectedValue:         expectedValue,
		errorStringContaining: errorStringContaining,
	}
}

type altFlagTestSimpleStringVar struct {
	displayName           string
	shortFlag             string
	usage                 string
	myVar                 string
	testClargs            []string
	expectedValue         *string
	errorStringContaining *string
}

// expectedErrorStringContaining implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) expectedErrorStringContaining() *string {
	return aft.errorStringContaining
}

// clargs implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) clargs() []string {
	return aft.testClargs
}

// flagSetName implements altFlagTest.
func (*altFlagTestSimpleStringVar) flagSetName() *string { return nil }

// setupFlagSet implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) setupFlagSet(f altflag.FlagSet) {
	f.StringVar(&aft.myVar, aft.displayName, aft.shortFlag, aft.usage)
}

// verify implements altFlagTest.
func (aft *altFlagTestSimpleStringVar) verify(t *testing.T, f altflag.FlagSet) {
	assert.Equal(t, aft.myVar, *aft.expectedValue)
}

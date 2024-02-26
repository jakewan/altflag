package altflag_test

import (
	"testing"

	"github.com/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

func newAltFlagTestCountVar(
	displayName string,
	shortFlag string,
	usage string,
	maxAllowed int,
	clargs []string,
	expectedCount int,
	setupErrorStringContaining *string,
	parseErrorStringContaining *string) altFlagTest {
	return &altFlagTestCountVar{
		displayName:                displayName,
		shortFlag:                  shortFlag,
		usage:                      usage,
		maxAllowed:                 maxAllowed,
		testClargs:                 clargs,
		expectedCount:              expectedCount,
		setupErrorStringContaining: setupErrorStringContaining,
		parseErrorStringContaining: parseErrorStringContaining,
	}
}

type altFlagTestCountVar struct {
	displayName                string
	shortFlag                  string
	usage                      string
	maxAllowed                 int
	myVar                      int
	testClargs                 []string
	expectedCount              int
	setupErrorStringContaining *string
	parseErrorStringContaining *string
}

// clargs implements altFlagTest.
func (aft *altFlagTestCountVar) clargs() []string {
	return aft.testClargs
}

// expectedParseErrorStringContaining implements altFlagTest.
func (aft *altFlagTestCountVar) expectedParseErrorStringContaining() *string {
	return aft.parseErrorStringContaining
}

// expectedSetupErrorStringContaining implements altFlagTest.
func (aft *altFlagTestCountVar) expectedSetupErrorStringContaining() *string {
	return aft.setupErrorStringContaining
}

// flagSetName implements altFlagTest.
func (*altFlagTestCountVar) flagSetName() *string { return nil }

// setupFlagSet implements altFlagTest.
func (aft *altFlagTestCountVar) setupFlagSet(f altflag.FlagSet) error {
	_, err := f.CountVar(&aft.myVar, aft.displayName, aft.shortFlag, aft.usage, aft.maxAllowed)
	return err
}

// verify implements altFlagTest.
func (aft *altFlagTestCountVar) verify(t *testing.T, f altflag.FlagSet) {
	assert.Equal(t, aft.expectedCount, aft.myVar)
}

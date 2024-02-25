package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

func newAltFlagTestMultiStringVar(
	displayNames []string,
	shortFlags []string,
	usages []string,
	clargs []string,
	expectedValues []*string,
	setupErrorStringContaining *string,
	parseErrorStringContaining *string) altFlagTest {
	vars := []*string{}
	for range len(displayNames) {
		vars = append(vars, ptr(""))
	}
	return &altFlagTestMultiStringVar{
		displayNames:               displayNames,
		shortFlags:                 shortFlags,
		usages:                     usages,
		myVars:                     vars,
		testClargs:                 clargs,
		expectedValues:             expectedValues,
		setupErrorStringContaining: setupErrorStringContaining,
		parseErrorStringContaining: parseErrorStringContaining,
	}
}

type altFlagTestMultiStringVar struct {
	displayNames               []string
	shortFlags                 []string
	usages                     []string
	myVars                     []*string
	testClargs                 []string
	expectedValues             []*string
	setupErrorStringContaining *string
	parseErrorStringContaining *string
}

// expectedSetupErrorStringContaining implements altFlagTest.
func (aft *altFlagTestMultiStringVar) expectedSetupErrorStringContaining() *string {
	return aft.setupErrorStringContaining
}

// clargs implements altFlagTest.
func (aft *altFlagTestMultiStringVar) clargs() []string {
	return aft.testClargs
}

// expectedParseErrorStringContaining implements altFlagTest.
func (aft *altFlagTestMultiStringVar) expectedParseErrorStringContaining() *string {
	return aft.parseErrorStringContaining
}

// flagSetName implements altFlagTest.
func (*altFlagTestMultiStringVar) flagSetName() *string { return nil }

// setupFlagSet implements altFlagTest.
func (aft *altFlagTestMultiStringVar) setupFlagSet(f altflag.FlagSet) error {
	for i := range len(aft.displayNames) {
		if _, err := f.StringVar(aft.myVars[i], aft.displayNames[i], aft.shortFlags[i], aft.usages[i]); err != nil {
			return err
		}
	}
	return nil
}

// verify implements altFlagTest.
func (aft *altFlagTestMultiStringVar) verify(t *testing.T, f altflag.FlagSet) {
	for i := range len(aft.displayNames) {
		assert.Equal(t, *aft.expectedValues[i], *aft.myVars[i])
	}
}

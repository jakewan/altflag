package altflag_test

import (
	"testing"

	"github.com/cbsinteractive/jakewan/altflag"
	"github.com/stretchr/testify/assert"
)

func newAltFlagTestNoop() altFlagTest { return &altFlagTestNoop{} }

type altFlagTestNoop struct{}

// expectedSetupErrorStringContaining implements altFlagTest.
func (*altFlagTestNoop) expectedSetupErrorStringContaining() *string { return nil }

// expectedParseErrorStringContaining implements altFlagTest.
func (*altFlagTestNoop) expectedParseErrorStringContaining() *string { return nil }

// setupFlagSet implements altFlagTest.
func (*altFlagTestNoop) setupFlagSet(f altflag.FlagSet) error { return nil }

// clargs implements altFlagTest.
func (*altFlagTestNoop) clargs() []string { return []string{} }

// flagSetName implements altFlagTest.
func (*altFlagTestNoop) flagSetName() *string { return nil }

// verify implements altFlagTest.
func (*altFlagTestNoop) verify(t *testing.T, f altflag.FlagSet) {
	assert.Equal(t, "some-flagset", f.Name())
}

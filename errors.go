package altflag

import (
	"fmt"
	"slices"
	"strings"
)

func newErrEnvVarLookupMultipleMatch(givenArg string, matches []string) error {
	slices.Sort(matches)
	return errEnvVarLookupMultipleMatch{givenArg, matches}
}

type errEnvVarLookupMultipleMatch struct {
	givenArg string
	matches  []string
}

// Error implements error.
func (e errEnvVarLookupMultipleMatch) Error() string {
	return fmt.Sprintf("argument %s matched multiple flags: %s", e.givenArg, strings.Join(e.matches, ", "))
}

func newErrEnvVarLookupZeroMatch(givenArg string) error {
	return errEnvVarLookupZeroMatch{givenArg}
}

type errEnvVarLookupZeroMatch struct {
	givenArg string
}

// Error implements error.
func (e errEnvVarLookupZeroMatch) Error() string {
	return fmt.Sprintf("argument %s didn't match any known flags", e.givenArg)
}

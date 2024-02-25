package altflag

import "errors"

var errVarLookupZeroMatch = errors.New("zero matches")
var errVarLookupMultipleMatches = errors.New("multiple matches")

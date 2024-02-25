package altflag

import (
	"fmt"
	"strings"
)

type TargetVariable[T any] interface {
	DisplayName() string
	SetDefault(defaultValue T)
	SetValue(value T)
	Value() T
}

func newTargetVariable[T any](target *T, displayName string, shortFlag string, usage string) TargetVariable[T] {
	return &targetVariable[T]{
		displayName: displayName,
		shortFlag:   shortFlag,
		t:           target,
		usage:       usage,
	}
}

type targetVariable[T any] struct {
	displayName  string
	shortFlag    string
	t            *T
	defaultValue T
	usage        string
}

// DisplayName implements TargetVariable.
func (t *targetVariable[T]) DisplayName() string {
	return t.displayName
}

// SetDefault implements TargetVariable.
func (t *targetVariable[T]) SetDefault(defaultValue T) {
	t.defaultValue = defaultValue
}

// SetValue implements TargetVariable.
func (t *targetVariable[T]) SetValue(value T) {
	*t.t = value
}

// Value implements TargetVariable.
func (t *targetVariable[T]) Value() T {
	return *t.t
}

type FlagSet interface {
	Name() string
	Parse(args []string) error
	BoolVar(target *bool, displayName string, shortFlag string, usage string) (TargetVariable[bool], error)
	CountVar(target *int, displayName string, shortFlag string, usage string, maxAllowed int) (TargetVariable[int], error)
	StringVar(target *string, displayName string, shortFlag string, usage string) (TargetVariable[string], error)
}

func NewFlagSet(name string) FlagSet {
	return &flagSet{
		boolVars:   []TargetVariable[bool]{},
		name:       name,
		stringVars: []TargetVariable[string]{},
	}
}

type flagSet struct {
	boolVars   []TargetVariable[bool]
	countVars  []TargetVariable[int]
	name       string
	stringVars []TargetVariable[string]
}

// Name implements FlagSet.
func (f *flagSet) Name() string {
	return f.name
}

// Parse implements FlagSet.
func (f *flagSet) Parse(args []string) error {
	fmt.Printf("Args: %s\n", args)
	for _, v := range f.stringVars {
		fmt.Printf("Flag: %v\n", v)
	}
	for _, v := range f.boolVars {
		fmt.Printf("Flag: %v\n", v)
	}
	fmt.Printf("Arg length: %d\n", len(args))
	for currentArgIdx, currentArg := range args {
		if strings.HasPrefix(currentArg, "-") {
			// Currently looking at a flag.
			fmt.Printf("Current arg index: %d\n", currentArgIdx)
			fmt.Printf("Current arg: %s\n", currentArg)
			stringVar, err := f.findStringVar(currentArg)
			if err != nil {
				return err
			}
			if stringVar != nil {
				if len(args) > currentArgIdx+1 {
					nextArg := args[currentArgIdx+1]
					stringVar.SetValue(nextArg)
				}
				continue
			}
			countVar, err := f.findCountVar(currentArg)
			if err != nil {
				return err
			}
			if countVar != nil {
				countVar.SetValue(countVar.Value() + 1)
				continue
			}
		}
	}
	return nil
}

func (f *flagSet) assertSingleMatch(givenArg string) error {
	normalizedGivenArg := normalizeArgName(givenArg)
	matches := []string{}
	for _, v := range f.boolVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	for _, v := range f.countVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	for _, v := range f.stringVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	fmt.Printf("Matches: %v\n", matches)
	if len(matches) > 1 {
		return newErrEnvVarLookupMultipleMatch(givenArg, matches)
	}
	if len(matches) == 0 {
		return newErrEnvVarLookupZeroMatch(givenArg)
	}
	return nil
}

func (f *flagSet) assertZeroMatches(givenArg string) error {
	normalizedGivenArg := normalizeArgName(givenArg)
	matches := []string{}
	for _, v := range f.boolVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	for _, v := range f.countVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	for _, v := range f.stringVars {
		displayName := v.DisplayName()
		fmt.Printf("Current configured display name: %s\n", displayName)
		normalizedDisplayName := normalizeArgName(displayName)
		if strings.HasPrefix(normalizedDisplayName, normalizedGivenArg) {
			matches = append(matches, displayName)
		}
	}
	fmt.Printf("Matches: %v\n", matches)
	if len(matches) > 0 {
		return fmt.Errorf("argument %s is already configured", givenArg)
	}
	return nil
}

func (f *flagSet) findCountVar(givenArg string) (TargetVariable[int], error) {
	if err := f.assertSingleMatch(givenArg); err != nil {
		return nil, err
	}
	lowerCaseName := strings.ToLower(givenArg)
	for _, v := range f.countVars {
		currentNameLower := strings.ToLower(v.DisplayName())
		if strings.HasPrefix(currentNameLower, lowerCaseName) {
			return v, nil
		}
	}
	return nil, nil
}

func (f *flagSet) findStringVar(givenArg string) (TargetVariable[string], error) {
	if err := f.assertSingleMatch(givenArg); err != nil {
		return nil, err
	}
	lowerCaseName := strings.ToLower(givenArg)
	for _, v := range f.stringVars {
		currentNameLower := strings.ToLower(v.DisplayName())
		if strings.HasPrefix(currentNameLower, lowerCaseName) {
			return v, nil
		}
	}
	return nil, nil
}

// BoolVar implements FlagSet.
func (f *flagSet) BoolVar(target *bool, displayName string, shortFlag string, usage string) (TargetVariable[bool], error) {
	displayName = normalizeConfiguredDisplayName(displayName)
	if err := f.assertZeroMatches(displayName); err != nil {
		return nil, err
	}
	t := newTargetVariable[bool](target, displayName, shortFlag, usage)
	f.boolVars = append(f.boolVars, t)
	return t, nil
}

// CountVar implements FlagSet.
func (f *flagSet) CountVar(target *int, displayName string, shortFlag string, usage string, maxAllowed int) (TargetVariable[int], error) {
	displayName = normalizeConfiguredDisplayName(displayName)
	if err := f.assertZeroMatches(displayName); err != nil {
		return nil, err
	}
	t := newTargetVariable[int](target, displayName, shortFlag, usage)
	f.countVars = append(f.countVars, t)
	return t, nil
}

// StringVar implements FlagSet.
func (f *flagSet) StringVar(target *string, displayName string, shortFlag string, usage string) (TargetVariable[string], error) {
	displayName = normalizeConfiguredDisplayName(displayName)
	if err := f.assertZeroMatches(displayName); err != nil {
		return nil, err
	}
	t := newTargetVariable[string](target, displayName, shortFlag, usage)
	f.stringVars = append(f.stringVars, t)
	return t, nil
}

func normalizeArgName(givenArg string) string {
	return strings.ToLower(strings.TrimLeft(givenArg, "-"))
}

func normalizeConfiguredDisplayName(displayName string) string {
	displayName = strings.TrimLeft(displayName, "-")
	for {
		if strings.HasPrefix(displayName, "--") {
			break
		}
		displayName = fmt.Sprintf("-%s", displayName)
	}
	return displayName
}

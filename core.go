package xscope

import "strings"

const (
	Delimeter = " "
)

type Scope string

func (s Scope) Slice() ScopeSlice {
	scopeMap := map[string]struct{}{}
	for _, scope := range strings.Fields(string(s)) {
		scopeMap[scope] = struct{}{}
	}

	uniqueScopes := make([]string, len(scopeMap))
	i := 0
	for scope := range scopeMap {
		uniqueScopes[i] = scope
		i += 1
	}

	return uniqueScopes
}

type ScopeSlice []string

func (s ScopeSlice) Len() int {
	return len(s)
}

func (s ScopeSlice) SqlPlaceholdler() string {
	placeholders := make([]string, len(s))
	for i := range s {
		placeholders[i] = "?"
	}

	return strings.Join(placeholders, ",")
}

func (s ScopeSlice) String() string {
	return strings.Join(s, Delimeter)
}

func (s ScopeSlice) Has(scope string) bool {
	for _, t := range s {
		if t == scope {
			return true
		}
	}

	return false
}

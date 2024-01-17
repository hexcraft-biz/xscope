package xscope

import "strings"

const (
	Delimeter = " "
)

type String string

func (s String) Slice() Slice {
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

type Slice []string

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) SqlPlaceholdler() string {
	placeholders := make([]string, len(s))
	for i := range s {
		placeholders[i] = "?"
	}

	return strings.Join(placeholders, ",")
}

func (s Slice) String() String {
	return String(strings.Join(s, Delimeter))
}

func (s Slice) Has(scope string) bool {
	for _, t := range s {
		if t == scope {
			return true
		}
	}

	return false
}

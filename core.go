package xscope

import (
	"fmt"
	"strings"
)

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

func (s Slice) AnySlice() []any {
	result := make([]any, len(s))
	for i, t := range s {
		result[i] = t
	}

	return result
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) GetNamedPlaceholder() string {
	placeholders := make([]string, len(s))
	for i := range s {
		placeholders[i] = fmt.Sprintf(":scope_%d", i)
	}

	return strings.Join(placeholders, ",")
}

func (s Slice) GetVarPlaceholder() string {
	placeholders := make([]string, len(s))
	for i := range s {
		placeholders[i] = "?"
	}

	return strings.Join(placeholders, ",")
}

func (s Slice) GetNamedArgs() map[string]any {
	args := map[string]any{}
	for i, t := range s {
		args[fmt.Sprintf("scope_%d", i)] = t
	}

	return args
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

package util

import "strings"

func IsUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func IsLower(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func IsDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func IsLetter(s rune) bool {
	return IsUpper(s) || IsLower(s)
}

func ToUpper(s rune) rune {
	if IsLower(s) {
		return s - 'a' + 'A'
	}
	return s
}

func ToLower(s rune) rune {
	if IsUpper(s) {
		return s - 'A' + 'a'
	}
	return s
}

func IsAllUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c == '_' {
			continue
		}
		if !IsUpper(c) && !IsDigit(c) {
			return false
		}
	}
	return true
}

func IsAllLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c == '_' {
			continue
		}
		if !IsLower(c) && !IsDigit(c) {
			return false
		}
	}
	return true
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c == '_' {
			continue
		}
		return IsUpper(c)
	}
	return false
}

func IsUncapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c == '_' {
			continue
		}
		return IsLower(c)
	}
	return false
}

func ToPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	prefix, s := CutPrefixUnderscore(s)
	s, suffix := CutSuffixUnderscore(s)
	isAllUpper := IsAllUpper(s)
	isUpper := false // deal with CAPITALS_WITH_UNDERSCORES case
	doCapitalize := true
	ret := ""
	s = strings.TrimPrefix(s, "_")
	for _, c := range s {
		// skip all _ (hello___world -> helloWorld)
		if c == '_' {
			doCapitalize = true
			continue
		}
		// found _ before || is first char (hello_world -> HelloWorld)
		if doCapitalize {
			ret += string(ToUpper(c))
			doCapitalize = false
			isUpper = IsUpper(c)
			continue
		}
		// found UPPER char before (HELLO -> Hello)
		if isAllUpper && isUpper {
			ret += string(ToLower(c))
			doCapitalize = false
			isUpper = IsUpper(c)
			continue
		}
		ret += string(c)
		doCapitalize = false
		isUpper = IsUpper(c)
	}
	ret = prefix + ret + suffix
	return ret
}

func TocamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	prefix, s := CutPrefixUnderscore(s)
	s, suffix := CutSuffixUnderscore(s)
	isAllUpper := IsAllUpper(s)
	isUpper := false // deal with CAPITALS_WITH_UNDERSCORES case
	doCapitalize := false
	ret := ""
	for i, c := range s {
		// skip all _ (hello___world -> helloWorld)
		if c == '_' {
			doCapitalize = true
			continue
		}
		// lower first char (HelloWorld -> helloWorld)
		if i == 0 {
			ret += string(ToLower(c))
			isUpper = IsUpper(c)
			continue
		}
		// found _ before (hello_world -> helloWorld)
		if doCapitalize {
			ret += string(ToUpper(c))
			doCapitalize = false
			isUpper = IsUpper(c)
			continue
		}
		// found UPPER char before (HELLO -> Hello)
		if isAllUpper && isUpper {
			ret += string(ToLower(c))
			doCapitalize = false
			isUpper = IsUpper(c)
			continue
		}
		ret += string(c)
		doCapitalize = false
		isUpper = IsUpper(c)
	}
	ret = prefix + ret + suffix
	return ret
}

func Tosnake_case(s string) string {
	if len(s) == 0 {
		return s
	}
	prefix, s := CutPrefixUnderscore(s)
	s, suffix := CutSuffixUnderscore(s)
	isAllUpper := IsAllUpper(s)
	isUpper := false // deal with CAPITALS_WITH_UNDERSCORES case
	hasUnderScore := false
	ret := ""
	for i, c := range s {
		if c == '_' {
			// skip multiple _ (hello___world -> hello_world)
			if !hasUnderScore {
				ret += string(c)
				hasUnderScore = true
			}
			continue
		}
		// lower first char (HelloWorld -> hello_world) (avoid Hello -> _hello)
		if i == 0 {
			ret += string(ToLower(c))
			isUpper = IsUpper(c)
			hasUnderScore = false
			continue
		}
		if IsUpper(c) {
			// skip multiple _ (hello_World -> hello_world) (avoid hello_World -> hello__world)
			// skip multiple upper (HELLO -> hello) (avoid HELLO -> h_e_l_l_o)
			if !hasUnderScore && !(isAllUpper && isUpper) {
				ret += "_"
				hasUnderScore = true
			}
			ret += string(ToLower(c))
			isUpper = true
			hasUnderScore = false
			continue
		}
		ret += string(c)
		isUpper = false
		hasUnderScore = false
	}
	ret = prefix + ret + suffix
	return ret
}

func ToALLCAP_CASE(s string) string {
	if len(s) == 0 {
		return s
	}
	prefix, s := CutPrefixUnderscore(s)
	s, suffix := CutSuffixUnderscore(s)
	isAllUpper := IsAllUpper(s)
	isUpper := false // deal with CAPITALS_WITH_UNDERSCORES case
	hasUnderScore := false
	ret := ""
	for i, c := range s {
		if c == '_' {
			// skip multiple _ (hello___world -> HELLO_WORLD)
			if !hasUnderScore {
				ret += string(c)
				hasUnderScore = true
			}
			continue
		}
		// upper first char (HelloWorld -> HELLO_WORLD) (avoid Hello -> _HELLO)
		if i == 0 {
			ret += string(ToUpper(c))
			isUpper = IsUpper(c)
			hasUnderScore = false
			continue
		}
		if IsUpper(c) {
			// skip multiple _ (hello_World -> HELLO_WORLD) (avoid hello_World -> HELLO__WORLD)
			// skip multiple upper (HELLO -> HELLO) (avoid HELLO -> H_E_L_L_O)
			if !hasUnderScore && !(isAllUpper && isUpper) {
				ret += "_"
				hasUnderScore = true
			}
			ret += string(ToUpper(c))
			isUpper = true
			hasUnderScore = false
			continue
		}
		ret += string(ToUpper(c))
		isUpper = false
		hasUnderScore = false
	}
	ret = prefix + ret + suffix
	return ret
}

func CutPrefixUnderscore(s string) (prefix string, rest string) {
	if len(s) == 0 {
		return "", ""
	}
	if s[0] != '_' {
		return "", s
	}
	for i, c := range s {
		if c != '_' {
			return s[:i], s[i:]
		}
	}
	return s, ""
}

func CutSuffixUnderscore(s string) (rest string, suffix string) {
	if len(s) == 0 {
		return "", ""
	}
	if s[len(s)-1] != '_' {
		return s, ""
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '_' {
			return s[:i+1], s[i+1:]
		}
	}
	return "", s
}

package util

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
	isFirst := true
	isUpper := false
	doCapitalize := true
	ret := ""
	for _, c := range s {
		if c == '_' {
			if isFirst {
				ret += string(c)
				continue
			}
			doCapitalize = true
			continue
		}
		if doCapitalize {
			ret += string(ToUpper(c))
			doCapitalize = false
			isFirst = false
			isUpper = IsUpper(c)
			continue
		}
		if isUpper {
			ret += string(ToLower(c))
			isUpper = IsUpper(c)
			continue
		}
		ret += string(c)
		isUpper = IsUpper(c)
	}
	return ret
}

func TocamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	isFirst := true
	isUpper := false
	doCapitalize := false
	ret := ""
	for _, c := range s {
		if c == '_' {
			if isFirst {
				ret += string(c)
				continue
			}
			doCapitalize = true
			continue
		}
		if isFirst {
			ret += string(ToLower(c))
			doCapitalize = false
			isFirst = false
			isUpper = IsUpper(c)
			continue
		}
		if doCapitalize {
			ret += string(ToUpper(c))
			doCapitalize = false
			isUpper = IsUpper(c)
			continue
		}
		if isUpper {
			ret += string(ToLower(c))
			isUpper = IsUpper(c)
			continue
		}
		ret += string(c)
		isUpper = IsUpper(c)
	}
	return ret
}

func Tosnake_case(s string) string {
	if len(s) == 0 {
		return s
	}
	isFirst := true
	isUpper := false
	hasUnderScore := false
	ret := ""
	for _, c := range s {
		if c == '_' {
			if isFirst {
				ret += string(c)
				continue
			}
			if !hasUnderScore {
				ret += string(c)
				hasUnderScore = true
			}
			continue
		}
		if IsUpper(c) {
			if !isFirst && !hasUnderScore && !isUpper {
				ret += "_"
				hasUnderScore = true
			}
			ret += string(ToLower(c))
			isFirst = false
			isUpper = true
			hasUnderScore = false
			continue
		}
		ret += string(c)
		isFirst = false
		isUpper = false
		hasUnderScore = false
	}
	return ret
}

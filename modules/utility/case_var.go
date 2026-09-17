package utility

import (
	"strings"
)

// CaseVarEngine provides case variation utilities.
type CaseVarEngine struct{}

// NewCaseVarEngine creates a new CaseVarEngine.
func NewCaseVarEngine() *CaseVarEngine {
	return &CaseVarEngine{}
}

// RandomCase applies random casing to a string.
func (e *CaseVarEngine) RandomCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i%2 == 0 {
			result.WriteString(strings.ToUpper(string(r)))
		} else {
			result.WriteString(strings.ToLower(string(r)))
		}
	}
	return result.String()
}

// AlternatingCase applies alternating upper/lower case.
func (e *CaseVarEngine) AlternatingCase(s string) string {
	var result strings.Builder
	upper := true
	for _, r := range s {
		if upper {
			result.WriteString(strings.ToUpper(string(r)))
		} else {
			result.WriteString(strings.ToLower(string(r)))
		}
		upper = !upper
	}
	return result.String()
}

// TitleCaseEveryWord capitalizes the first letter of each word.
func (e *CaseVarEngine) TitleCaseEveryWord(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// SnakeToCamel converts snake_case to camelCase.
func (e *CaseVarEngine) SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(string(parts[i][0])) + strings.ToLower(parts[i][1:])
		}
	}
	return strings.Join(parts, "")
}

// CamelToSnake converts camelCase to snake_case.
func (e *CaseVarEngine) CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteString("_")
			result.WriteString(strings.ToLower(string(r)))
		} else {
			result.WriteString(string(r))
		}
	}
	return result.String()
}

// LeetSpeak converts letters to leet speak equivalents.
func (e *CaseVarEngine) LeetSpeak(s string) string {
	replacements := map[rune]string{
		'a': "4", 'A': "4",
		'e': "3", 'E': "3",
		'i': "1", 'I': "1",
		'o': "0", 'O': "0",
		's': "5", 'S': "5",
		't': "7", 'T': "7",
		'l': "1", 'L': "1",
		'z': "2", 'Z': "2",
	}

	var result strings.Builder
	for _, r := range s {
		if repl, ok := replacements[r]; ok {
			result.WriteString(repl)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

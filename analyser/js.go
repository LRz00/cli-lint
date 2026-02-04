package analyser

import (
	"regexp"
	"strings"
)

func UsesVar(jsCode string) bool {
	return strings.Contains(jsCode, "var ")
}

func UsesWeakComparisons(jsCode string) bool {
	return strings.Contains(jsCode, "==") || strings.Contains(jsCode, "!=")
}

type SemiStats struct {
	Lines int
	Semis int
}

func AnalyzeSemi(code string) SemiStats {
	lines := strings.Split(code, "\n")
	return SemiStats{
		Lines: len(lines),
		Semis: strings.Count(code, ";"),
	}
}

type QuoteStats struct {
	Single int
	Double int
}

func AnalyzeQuotes(code string) QuoteStats {
	return QuoteStats{
		Single: strings.Count(code, "'"),
		Double: strings.Count(code, "\""),
	}
}

func UsesConsole(code string) bool {
	return strings.Contains(code, "console.")
}

var varDeclRegex = regexp.MustCompile(`\b(const|let)\s+([a-zA-Z_][a-zA-Z0-9_]*)`)

func CountUnusedVars(code string) int {
	matches := varDeclRegex.FindAllStringSubmatch(code, -1)
	unused := 0

	for _, m := range matches {
		varName := m[2]
		count := strings.Count(code, varName)

		if count == 1 {
			unused++
		}
	}

	return unused
}

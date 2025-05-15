package filterpipe

import (
	"strings"
	"unicode"
)

type FilterPipe struct {
}

func New() *FilterPipe {
	return &FilterPipe{}
}

func (f *FilterPipe) Fileter(str string) string {
	replacer := strings.NewReplacer(
		"0", "o", "1", "l", "3", "e", "4", "a", "5", "s", "7", "t", "8", "b", "9", "g",
		"а", "a", "А", "a", // кириллица → латиница
		"в", "b", "В", "b",
		"с", "c", "С", "c",
		"е", "e", "Е", "e",
		"н", "h", "Н", "h",
		"к", "k", "К", "k",
		"м", "m", "М", "m",
		"о", "o", "О", "o",
		"р", "p", "Р", "p",
		"т", "t", "Т", "t",
		"х", "x", "Х", "x",
		"у", "y", "У", "y",
		"л", "l", "Л", "l",
		"г", "g", "Г", "g",
		"ш", "w", "Ш", "w",
	)

	normalized := strings.ToLower(str)

	normalized = replacer.Replace(normalized)

	normalized = strings.ReplaceAll(normalized, " ", "")

	var filtered strings.Builder
	for _, r := range normalized {
		if (unicode.IsLetter(r) && r <= unicode.MaxASCII) || unicode.IsDigit(r) || r == '_' {
			filtered.WriteRune(r)
		}
	}

	return filtered.String()
}

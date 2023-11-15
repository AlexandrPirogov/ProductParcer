// package parser holds table of avaible parsers for parsing
package parser

import (
	"log"
	"product-parser/comparator"
	"product-parser/parser/csv"
	"product-parser/parser/json"
)

type Extension string

const (
	JSON Extension = ".json"
	CSV  Extension = ".csv"
)

type Parser[T any] interface {
	// Parses given text and unmarshal into T instance
	Parse(text string) (T, error)
}

// New return instance of parser depending on given file extension
//
// Pre-cond: given extension of file
func New[T comparator.Comparable[T]](ext string) Parser[T] {
	avaibleParsers := parsers[T]()
	parser, ok := avaibleParsers[Extension(ext)]
	if !ok {
		log.Fatalf("given unsupported extension %s", ext)
	}

	return parser
}

func parsers[T comparator.Comparable[T]]() map[Extension]Parser[T] {
	return map[Extension]Parser[T]{
		CSV:  csv.New[T](),
		JSON: json.New[T](),
	}
}

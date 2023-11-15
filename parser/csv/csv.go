// package json holds CSV parser
package csv

import (
	"fmt"

	"github.com/jszwec/csvutil"
)

// New returns pointer to the new instance of CSV Parser
func New[T any]() *CSV[T] {
	return &CSV[T]{}
}

type CSV[T any] struct {
	Header []byte
}

// Parse given string and unmarshal into given T instance
// Firstly it parses Header of CSV file and returns error "header was created"
// Then Parse Unmarshal given text to T instance
//
// Pre-cond: given string
//
// Post-cond: unmarshal given text. Return instance and error status
func (c *CSV[T]) Parse(text string) (T, error) {

	var fail T
	var success []T
	if c.Header == nil {
		c.Header = []byte(text)
		return fail, fmt.Errorf("header was created")
	}

	toUnmarshal := c.buildCSVInput(text)
	err := csvutil.Unmarshal(toUnmarshal, &success)
	if err != nil {
		return fail, err
	}
	return success[0], nil
}

func (c *CSV[T]) buildCSVInput(text string) []byte {
	toUnmarshal := make([]byte, 0, len(c.Header)+len(text)+1)
	toUnmarshal = append(toUnmarshal, c.Header...)
	toUnmarshal = append(toUnmarshal, '\n')
	toUnmarshal = append(toUnmarshal, []byte(text)...)
	return toUnmarshal
}

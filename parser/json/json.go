// package json holds JSON parser
package json

import (
	"encoding/json"
)

// New returns pointer to the new instance of JSON Parser
func New[T any]() *JSON[T] {
	return &JSON[T]{}
}

type JSON[T any] struct{}

// Parse given string and unmarshal into given T instance
//
// Pre-cond: given string
//
// Post-cond: unmarshal given text. Return instance and error status
func (c *JSON[T]) Parse(text string) (T, error) {
	bytes := []byte(text)
	if bytes[len(bytes)-1] == ',' {
		bytes = bytes[:len(bytes)-1]
	}

	var res T
	err := json.Unmarshal(bytes, &res)
	return res, err
}

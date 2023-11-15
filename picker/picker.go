// package picker contains tools for working with files
package picker

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Ext  string
	FD   *os.File
}

// Open opens file by giving path and returns File struct
//
// Pre-cond: given path
//
// Post-cond: if file exists -> returns File struct instance
func Open(path string) File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	f := File{
		Name: filepath.Base(path),
		Ext:  filepath.Ext(path),
		FD:   file,
	}

	return f
}

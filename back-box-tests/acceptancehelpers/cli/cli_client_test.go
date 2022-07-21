package cli_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestSomething(t *testing.T) {
	//is := is.New(t)

	dir := "/Users/Michael/Documents/training/temperature-converter/cmd"
	build := "converter"

	file := filepath.Join(dir, build)
	fmt.Println(file)

	fmt.Println("STARTTTTTTTTTTTTTT " + file + " ENNNNNNNNNNNNDDDDD")
}

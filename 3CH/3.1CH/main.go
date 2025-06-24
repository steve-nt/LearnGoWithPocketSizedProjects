package main

import (
	"fmt"
	"os"
	"github.com/steve-nt/bookworm-example/bookworm"
)

func main() {
	bookworms, err := bookworm.LoadBookworms (
		"testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr,
			"failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}

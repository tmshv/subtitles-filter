package main

import (
	"fmt"
	"os"

	f "github.com/tmshv/subtitles-filter/internal/filter"
	"github.com/tmshv/subtitles-filter/internal/vtt"
)

var filter f.Filter = f.All(
	f.MinLen(2),
	f.Not(f.Any(
		f.Eq("."),
		f.Substr("DimaTorzok"),
	)),
)

func run(filename string) error {
	s, err := vtt.OpenFile(filename)
	if err != nil {
		return err
	}

	// Write head
	fmt.Printf("%s\n\n", s.Head())

	// Write filterd records
	for rec := range s.Iter() {
		if filter.Test(rec.Text()) {
			fmt.Printf("%s\n%s\n\n", rec.Time(), rec.Text())
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 1 {
		panic("no filename")
	}
	var filename string = os.Args[1]
	err := run(filename)
	if err != nil {
		panic(err)
	}
}

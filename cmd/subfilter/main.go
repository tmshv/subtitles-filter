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

func run() error {
	var sub *vtt.VTT
	var err error
	if len(os.Args) <= 1 {
		sub, err = vtt.Scan(os.Stdin)
	} else {
		var filename string = os.Args[1]
		sub, err = vtt.OpenFile(filename)
	}
	if err != nil {
		return err
	}

	// Write head
	fmt.Printf("%s\n\n", sub.Head())

	// Write filterd records
	for rec := range sub.Iter() {
		if filter.Test(rec.Text()) {
			fmt.Printf("%s\n%s\n\n", rec.Time(), rec.Text())
		}
	}

	return nil
}

func main() {
    if err := run(); err != nil {
		panic(err)
	}
}

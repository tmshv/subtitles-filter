package filter

type prev struct {
	prev string
}

func (f *prev) Test(text string) bool {
	defer f.save(text)
	return text != f.prev
}

func (f *prev) save(prev string) {
	f.prev = prev
}

func Prev() *prev {
	return &prev{}
}

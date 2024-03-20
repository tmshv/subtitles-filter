package filter

type eq struct {
	val string
}

func (f *eq) Test(text string) bool {
	return text == f.val
}

type not struct {
	child Filter
}

func (f *not) Test(text string) bool {
	return !f.child.Test(text)
}

type all []Filter

func (f all) Test(text string) bool {
	for _, x := range f {
		if !x.Test(text) {
			return false
		}
	}
	return true
}

type any []Filter

func (f any) Test(text string) bool {
	for _, x := range f {
		if x.Test(text) {
			return true
		}
	}
	return false
}

func Eq(val string) Filter {
	return &eq{val}
}

func Not(f Filter) Filter {
	return &not{f}
}

func All(fs ...Filter) Filter {
	return all(fs)
}

func Any(fs ...Filter) Filter {
	return any(fs)
}

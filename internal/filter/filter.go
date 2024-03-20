package filter

type Filter interface {
	Test(string) bool
}


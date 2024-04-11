package filter

import "testing"

func TestPrev(t *testing.T) {
    f := Prev()
    sample := "subtitle"

    if !f.Test(sample) {
        t.Error("Prev should pass first call")
    }
    if f.Test(sample) {
        t.Error("Prev should not pass second call")
    }
    if !f.Test("next") {
        t.Error("Prev should pass if text changed")
    }
}

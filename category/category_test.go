package category

import (
	"strings"
	"testing"
)

func TestParseResponse(t *testing.T) {
	r, err := ParseResponse(strings.NewReader("{}"))
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Categories) != 0 {
		t.Fatal()
	}

	r, err = ParseResponse(strings.NewReader(`{"categories": [{"id": 1}, {"id": 2}]}`))
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Categories) != 2 {
		t.Fatal()
	}

	println(r)
}

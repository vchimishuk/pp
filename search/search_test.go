package search

import (
	"testing"

	"github.com/vchimishuk/pp/category"
)

func TestSearch(t *testing.T) {
	cats := []*category.SourceCategory{
		&category.SourceCategory{Id: 1, Parent: 0},
		&category.SourceCategory{Id: 2, Parent: 0},
		&category.SourceCategory{Id: 3, Parent: 2},
		&category.SourceCategory{Id: 4, Parent: 2},
		&category.SourceCategory{Id: 5, Parent: 4},
		&category.SourceCategory{Id: 6, Parent: 5},
		&category.SourceCategory{Id: 7, Parent: 5}}
	roots := Tree(cats, 0)
	assert(t, 2, len(roots))
	assert(t, 1, roots[0].Id)
	assert(t, 2, roots[1].Id)

	roots = Tree(cats, -1)
	assert(t, 2, len(roots))
	assert(t, 1, roots[0].Id)
	assert(t, 2, roots[1].Id)

	roots = Tree(cats, 1)
	assert(t, 2, len(roots))
	assert(t, 1, roots[0].Id)
	assert(t, 2, roots[1].Id)

	roots = Tree(cats, 3)
	assert(t, 2, len(roots))
	assert(t, 2, roots[0].Id)
	assert(t, 1, roots[1].Id)
	assert(t, 2, len(roots[0].Categories))
	assert(t, 3, roots[0].Categories[0].Id)
	assert(t, 0, len(roots[0].Categories[0].Categories))
	assert(t, 4, roots[0].Categories[1].Id)
	assert(t, 0, len(roots[0].Categories[1].Categories))

	roots = Tree(cats, 6)
	assert(t, 2, len(roots))
	assert(t, 2, roots[0].Id)
	assert(t, 1, roots[1].Id)
	assert(t, 1, len(roots[0].Categories))
	assert(t, 0, len(roots[1].Categories))
	assert(t, 4, roots[0].Categories[0].Id)
	assert(t, 1, len(roots[0].Categories[0].Categories))
	assert(t, 5, roots[0].Categories[0].Categories[0].Id)
	assert(t, 2, len(roots[0].Categories[0].Categories[0].Categories))
	assert(t, 6, roots[0].Categories[0].Categories[0].Categories[0].Id)
	assert(t, 7, roots[0].Categories[0].Categories[0].Categories[1].Id)
}

func assert(t *testing.T, exp, actual int) {
	if exp != actual {
		t.Fatalf("Expected %d but %d given", exp, actual)
	}
}

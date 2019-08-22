package search

import (
	"github.com/vchimishuk/pp/category"
)

type ClientCategory struct {
	category.SourceCategory
	Categories []*ClientCategory `json:"categories"`
}

func Tree(cats []*category.SourceCategory, id int) []*ClientCategory {
	// Parent to children mapping index.
	parentIdx := map[int][]*category.SourceCategory{}
	for _, c := range cats {
		parentIdx[c.Parent] = append(parentIdx[c.Parent], c)
	}

	// Id to category mapping index.
	idIdx := map[int]*category.SourceCategory{}
	for _, c := range cats {
		idIdx[c.Id] = c
	}

	// Build tree containing id-node.
	sc, ok := idIdx[id]
	var c *ClientCategory
	if sc != nil {
		c = &ClientCategory{SourceCategory: *sc}
	}
	for ok {
		sp, ok := idIdx[c.Parent]
		if !ok {
			break
		}
		p := &ClientCategory{SourceCategory: *sp}
		p.Categories = []*ClientCategory{c}
		// Add siblings.
		if c.Id == id {
			for _, ch := range parentIdx[p.Id] {
				if ch.Id != id {
					cl := &ClientCategory{SourceCategory: *ch}
					p.Categories = append(p.Categories, cl)
				}
			}
		}
		c = p
	}

	roots := make([]*ClientCategory, 0, 3)
	if c != nil {
		roots = append(roots, c)
	}
	// Add other roots but self.
	for _, r := range parentIdx[0] {
		if c == nil || r.Id != c.Id {
			roots = append(roots,
				&ClientCategory{SourceCategory: *r})
		}
	}

	return roots
}

package trie

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestRuneTrie(t *testing.T) {

	r := NewRuneTrie()

	r.Put("123", 123)
	r.Put("12", 12)
	r.Put("13", 13)
	r.Put("25", 25)

	spew.Dump(r)

	assert.Equal(t, 25, r.Get("25"))

	r.Delete("25")
	assert.Equal(t, nil, r.Get("25"))
	spew.Dump(r)
}

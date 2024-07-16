package implementtrie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Insert(t *testing.T) {
	t.Run("search inserted", func(t *testing.T) {
		sut := Constructor()
		const s = "abc"

		sut.Insert(s)

		assert.True(t, sut.Search(s))
	})

	t.Run("search uninserted", func(t *testing.T) {
		sut := Constructor()
		const s = "abc"

		sut.Insert(s)

		assert.False(t, sut.Search("ijk"))
	})

	t.Run("searching by part of the word that has been inserted", func(t *testing.T) {
		sut := Constructor()
		const s = "abc"

		sut.Insert(s)

		assert.False(t, sut.Search("ab"))
	})

	t.Run("searching by prefix of inserted word", func(t *testing.T) {
		sut := Constructor()
		const s = "abc"

		sut.Insert(s)

		assert.True(t, sut.StartsWith("ab"))
	})

	t.Run("searching by prefix of uninserted word", func(t *testing.T) {
		sut := Constructor()
		const s = "abc"

		sut.Insert(s)

		assert.False(t, sut.StartsWith("ij"))
	})
}

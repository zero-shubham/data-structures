package trie_test

import (
	"math/rand"
	rand2 "math/rand/v2"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zero-shubham/data-structures/Trie/trie"
)

func randRange(min, max int) int {
	return rand2.IntN(max-min) + min
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int, sb *strings.Builder) string {
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func BenchmarkTrieReuseOverload(b *testing.B) {
	t := trie.NewTrie()

	b.Run("benchmark insert trie", func(b *testing.B) {
		sb := strings.Builder{}
		for i := 0; i < b.N; i++ {
			randStr := RandStringBytesMaskImprSrcSB(randRange(5, 500), &sb)

			err := t.Insert(randStr)
			if err != nil {
				b.Logf("generated string: %s", randStr)
			}
			assert.NoError(b, err)
			sb.Reset()
		}
	})

	randStrsTrie := make([]string, 0, 5000)
	sb := strings.Builder{}
	for i := 0; i < 5000; i++ {
		randStrsTrie = append(randStrsTrie, RandStringBytesMaskImprSrcSB(randRange(50, 500), &sb))
		t.Insert(randStrsTrie[i])
		sb.Reset()
	}

	b.Run("benchmark search trie", func(b *testing.B) {
		for i := 0; i < len(randStrsTrie); i++ {
			ok := t.Search(randStrsTrie[i])
			if !ok {
				b.Log("failed seach for: ", randStrsTrie[i])
			}
			assert.True(b, ok)
		}
	})

}

func BenchmarkTrieOptimized(b *testing.B) {
	t := trie.NewTrie()

	b.Run("benchmark insert trie", func(b *testing.B) {
		sb := strings.Builder{}
		for i := 0; i < b.N; i++ {
			randStr := RandStringBytesMaskImprSrcSB(randRange(5, 500), &sb)

			err := t.Insert(randStr)
			if err != nil {
				b.Logf("generated string: %s", randStr)
			}
			assert.NoError(b, err)

			sb.Reset()
		}
	})

	t = trie.NewTrie()
	randStrsTrie := make([]string, 0, 5000)
	sb := strings.Builder{}
	for i := 0; i < 5000; i++ {
		randStrsTrie = append(randStrsTrie, RandStringBytesMaskImprSrcSB(randRange(50, 500), &sb))
		t.Insert(randStrsTrie[i])
		sb.Reset()
	}

	b.Run("benchmark search trie", func(b *testing.B) {
		for i := 0; i < len(randStrsTrie); i++ {
			ok := t.Search(randStrsTrie[i])
			if !ok {
				b.Log("failed seach for: ", randStrsTrie[i])
			}
			assert.True(b, ok)
		}
	})
}

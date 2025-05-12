package triemap

import (
	"fmt"
	"strings"
)

type TrieMapNode struct {
	word     bool
	children map[rune]*TrieMapNode
}

func NewTrieMapNode() *TrieMapNode {
	return &TrieMapNode{
		word:     false,
		children: make(map[rune]*TrieMapNode, 36),
	}
}

type TrieMap struct {
	root *TrieMapNode
}

func NewTrieMap() *TrieMap {
	return &TrieMap{
		root: NewTrieMapNode(),
	}
}

func isUnsupportedChar(charInt int) bool {
	return (charInt > 57 && charInt < 97) || charInt > 122 || charInt < 48
}

func (t *TrieMap) Insert(word string) error {
	word = strings.ToLower(word)
	curr := t.root

	for idx, char := range word {
		charInt := int(char)

		if isUnsupportedChar(charInt) {
			return fmt.Errorf("unsupported character %c", char)
		}

		if _, ok := curr.children[char]; !ok {
			curr.children[char] = NewTrieMapNode()
		}

		if idx == len(word)-1 {
			curr.children[char].word = true
		}
		curr = curr.children[char]
	}

	return nil
}

func (t *TrieMap) Search(word string) bool {
	word = strings.ToLower(word)
	curr := t.root

	for _, char := range word {
		charInt := int(char)
		if isUnsupportedChar(charInt) {
			return false
		}

		if _, ok := curr.children[char]; !ok {
			return false
		}

		curr = curr.children[char]
	}

	return curr.word
}

func (t *TrieMap) FindPrefix(word string) bool {
	word = strings.ToLower(word)
	curr := t.root

	for _, char := range word {
		charInt := int(char)
		if isUnsupportedChar(charInt) {
			return false
		}

		if _, ok := curr.children[char]; !ok {
			return false
		}

		curr = curr.children[char]
	}

	return true
}

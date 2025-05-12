package trie

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	word     bool
	children []*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		word:     false,
		children: make([]*TrieNode, 36),
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func isUnsupportedChar(charInt int) bool {
	return (charInt > 57 && charInt < 97) || charInt > 122 || charInt < 48
}

func getIdxPos(charInt int) int {

	if charInt >= 48 && charInt <= 57 {
		return charInt - 48
	}

	// if charInt >= 97 && charInt <= 122 {
	// 	return charInt - 87
	// }

	return charInt - 87
}

func (t *Trie) Insert(word string) error {
	word = strings.ToLower(word)
	curr := t.root
	for _, char := range word {

		charInt := int(char)

		if isUnsupportedChar(charInt) {
			return fmt.Errorf("unsupported character %c", char)
		}

		idxPos := getIdxPos(charInt)

		if curr.children[idxPos] == nil {
			curr.children[idxPos] = NewTrieNode()
		}

		curr = curr.children[idxPos]
	}

	curr.word = true

	return nil
}

func (t *Trie) Search(word string) bool {
	word = strings.ToLower(word)
	curr := t.root

	for _, char := range word {
		charInt := int(char)

		if isUnsupportedChar(charInt) {
			return false
		}

		idxPos := getIdxPos(charInt)

		if curr.children[idxPos] == nil {
			return false
		}

		curr = curr.children[idxPos]
	}

	return curr.word
}

func (t *Trie) FindPrefix(word string) bool {
	word = strings.ToLower(word)
	curr := t.root

	for _, char := range word {
		charInt := int(char)

		if isUnsupportedChar(charInt) {
			return false
		}

		idxPos := getIdxPos(charInt)

		if curr.children[idxPos] == nil {
			return false
		}

		curr = curr.children[idxPos]
	}

	return true
}

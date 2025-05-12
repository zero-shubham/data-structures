package main

import (
	"fmt"

	"github.com/zero-shubham/data-structures/Trie/trie"
	triemap "github.com/zero-shubham/data-structures/Trie/trie_map"
)

func main() {
	t := trie.NewTrie()

	err := t.Insert("apple1290")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.Search("apple1290"))
	fmt.Println(t.Search("apple"))
	fmt.Println(t.FindPrefix("apple"))

	err = t.Insert("f1zgqe8fygvr8r0cp")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---------")

	tm := triemap.NewTrieMap()

	err = tm.Insert("apple1290")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tm.Search("apple1290"))
	fmt.Println(tm.Search("apple"))
	fmt.Println(tm.FindPrefix("apple"))
}

package trie

import "fmt"

// Trie specs
type Trie interface {
	HelloWorld()
}

type trie struct {
	value    rune
	children *[]trie
}

// New makes new trie instance
func New() Trie {
	return &trie{}
}

func (t *trie) HelloWorld() {
	fmt.Println("trie says hello world!")
}

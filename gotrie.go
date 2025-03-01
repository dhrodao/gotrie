package main

import (
	"fmt"
	"unicode/utf8"
)

// Structure which represents a node from the trie data structure
type Node struct {
	// Value of the current node
	Value rune
	// Whether the current node is a leaf
	IsLeaf bool
	// Children of the current node
	Children map[rune]*Node
}

// Structure which represents a trie
type Trie struct {
	Root *Node
}

// This function creates a new node
func NewNode(value rune) *Node {
	node := &Node{Value: value}

	node.Children = make(map[rune]*Node)

	return node
}

// This function creates a new trie
func NewTrie() *Trie {
	return &Trie{Root: NewNode('\000')}
}

// This function inserts a word into the trie
func (t *Trie) Insert(s string) {
	current := t.Root

	// for k, v := range s {
	// 	char := s[i]
	// 	if current.Children[char] == nil {
	// 		current.Children[char] = NewNode(char)
	// 	}
	// 	current = current.Children[char]
	// }

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		if current.Children[runeValue] == nil {
			current.Children[runeValue] = NewNode(runeValue)
		}

		current = current.Children[runeValue]
		w = width
	}

	current.IsLeaf = true
}

func (t *Trie) Delete(s string) {
	current := t.Root

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		if current == nil || current.Children[runeValue] == nil {
			return
		}

		current = current.Children[runeValue]
		w = width
	}

	current.IsLeaf = false
}

// This function searches for a word in the trie
func (t *Trie) Search(s string) bool {
	current := t.Root

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		if current == nil || current.Children[runeValue] == nil {
			return false
		}

		current = current.Children[runeValue]
		w = width
	}

	return current.IsLeaf
}

func main() {
	trie := NewTrie()

	word := "hello"
	trie.Insert(word)
	fmt.Printf("%s: %v\n", word, trie.Search(word))

	word = "hell"
	fmt.Printf("%s: %v\n", word, trie.Search(word))
	trie.Insert("hell")
	fmt.Printf("%s: %v\n", word, trie.Search(word))
	trie.Delete(word)
	fmt.Printf("%s: %v\n", word, trie.Search(word))

	word = "dhrodao@example.com"
	trie.Insert(word)
	fmt.Printf("%s: %v\n", word, trie.Search(word))
	word = "dhrodao@example."
	fmt.Printf("%s: %v\n", word, trie.Search(word))

	const r = "สวัสดี"
	trie.Insert(r)
	fmt.Printf("%s: %v\n", r, trie.Search(r))
}

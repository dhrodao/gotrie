package main

import "fmt"

// Structure which represents a node from the trie data structure
type Node struct {
	// Value of the current node
	Char byte
	// Whether the current node is a leaf
	IsLeaf bool
	// Children of the current node
	// Size is 128 because there are 128 ASCII characters
	Children [128]*Node
}

// Structure which represents a trie
type Trie struct {
	Root *Node
}

// This function creates a new node
func NewNode(char byte) *Node {
	node := &Node{Char: char}

	for i := 0; i < 128; i++ {
		node.Children[i] = nil
	}

	return node
}

// This function creates a new trie
func NewTrie() *Trie {
	return &Trie{Root: NewNode('\000')}
}

// This function inserts a word into the trie
func (t *Trie) Insert(word string) {
	current := t.Root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if current.Children[char] == nil {
			current.Children[char] = NewNode(char)
		}
		current = current.Children[char]
	}

	current.IsLeaf = true
}

func (t *Trie) Delete(word string) {
	current := t.Root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if current == nil || current.Children[char] == nil {
			return
		}
		current = current.Children[char]
	}

	current.IsLeaf = false
}

// This function searches for a word in the trie
func (t *Trie) Search(word string) bool {
	current := t.Root

	for i := 0; i < len(word); i++ {
		char := word[i]
		if current == nil || current.Children[char] == nil {
			return false
		}
		current = current.Children[char]
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
}

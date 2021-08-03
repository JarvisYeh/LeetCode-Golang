package main

type TrieNode208 struct {
	children []*TrieNode208
	isWord   bool
}

type Trie208 struct {
	root *TrieNode208
}

/** Initialize your data structure here. */
func Constructor() Trie208 {
	trie := Trie208{nil}
	trie.root = &TrieNode208{make([]*TrieNode208, 26), false}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie208) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		next := curr.children[word[i]-'a']
		if next == nil {
			next = &TrieNode208{make([]*TrieNode208, 26), false}
		}
		curr.children[word[i]-'a'] = next
		curr = next
	}
	if !curr.isWord {
		curr.isWord = true
	}
}

/** Returns if the word is in the trie. */
func (this *Trie208) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		next := curr.children[word[i]-'a']
		if next == nil {
			return false
		}
		curr = next
	}
	return curr.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie208) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		next := curr.children[prefix[i]-'a']
		if next == nil {
			return false
		}
		curr = next
	}
	return true
}

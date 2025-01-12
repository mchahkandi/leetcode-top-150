type TrieNode struct {
	Children map[string]*TrieNode
	EndOfWord bool
}

type Trie struct {
	Node *TrieNode
}


func Constructor() Trie {
	return Trie{Node: &TrieNode{ Children: map[string]*TrieNode{}, EndOfWord: false}}
}


func (this *Trie) Insert(word string)  {
	current := this.Node

	for _, ch := range word {
		if _, ok := current.Children[string(ch)]; !ok {
			current.Children[string(ch)] = &TrieNode{Children: map[string]*TrieNode{}, EndOfWord: false}
		}
		current = current.Children[string(ch)]
	}
	current.EndOfWord = true
}


func (this *Trie) Search(word string) bool {
	current := this.Node

	for _, ch := range word {
		if _, ok := current.Children[string(ch)]; !ok {
			return false
		}
		current = current.Children[string(ch)]
	}
	return current.EndOfWord
}


func (this *Trie) StartsWith(prefix string) bool {
	current := this.Node

	for _, ch := range prefix {
		if _, ok := current.Children[string(ch)]; !ok {
			return false
		}
		current = current.Children[string(ch)]
	}
	return true
}

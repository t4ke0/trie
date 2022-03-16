package main

import "fmt"

// Node
type Node struct {
	Children []*Node
	IsEnd    bool
}

// Insert
func Insert(w string, root *Node, index int) {
	if index == len(w) {
		root.IsEnd = true
		return
	}

	idx := uint8(w[index])

	if root == nil {
		root = new(Node)
	}

	if len(root.Children) == 0 {
		root.Children = make([]*Node, 256)
		root.Children[idx] = &Node{}
	} else {
		if root.Children[idx] == nil {
			root.Children[idx] = &Node{}
		}
	}

	index += 1
	Insert(w, root.Children[idx], index)

}

// Search
func (n *Node) Search(w string, index int, current *Node) {

	if current == nil {
		current = n
		n.Search(w, index, current)
		return
	}

	idx := uint8(w[index])

	if current.Children[idx] == nil {
		fmt.Println("not found")
		return
	}

	fmt.Println(string(w[index]))

	if current.Children[idx].IsEnd {
		fmt.Println("final found")
		return
	}

	index += 1
	n.Search(w, index, current.Children[idx])

}

// GetCompletion
func (n *Node) GetCompletion(w string) []string {

	var root *Node
	for _, c := range w {
		if root == nil {
			root = n
		}

		if len(root.Children) == 0 {
			fmt.Println("here")
			return []string{}
		}

		root = root.Children[uint8(c)]
	}

	out := []string{}

	getCases(root, "", &out)

	return out

}

func getCases(root *Node, w string, completions *[]string) {

	if root == nil {
		return
	}

	if root.IsEnd {
		*completions = append(*completions, w)
		return
	}

	if len(root.Children) == 0 {
		return
	}

	for c, child := range root.Children {
		if child != nil {
			w += string(rune(c))
			getCases(child, w, completions)
			w = ""
		}
	}
}

func main() {

	words := []string{
		"hello",
		"hammer",
		"other",
		"sun",
		"something",
	}

	root := new(Node)

	for _, word := range words {
		Insert(word, root, 0)
	}

	root.Search("hammer", 0, nil)

	fmt.Println(root.GetCompletion("s"))

}

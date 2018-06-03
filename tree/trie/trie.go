package trie

//Node a single node that composes the tree
type Node struct {
	value    string
	children map[string]*Node
}

//Trie tree
type Trie struct {
	root *Node
}

//New is create a new tree
func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
		},
	}
}

//Root is return a root
func (t *Trie) Root() *Node {
	return t.root
}

//Add is add a new nodo
func (t *Trie) Add(lCode []string, name string) *Node {
	node := t.root
	for i := range lCode {
		r := lCode[i]
		if n, ok := node.children[r]; ok {
			node = n
		} else {
			// create the
			node = node.NewChild(r, name)
		}
	}
	return node
}

//NewChild is create and init a child
func (n *Node) NewChild(key string, value string) *Node {
	node := &Node{
		value:    value,
		children: make(map[string]*Node),
	}
	n.children[key] = node
	return node
}

//Children is return a child
func (n Node) Children() map[string]*Node {
	return n.children
}

package gotreap

type Item interface {
	LessKey(*Item) bool
	LessPriority(*Item) bool
}

type Treap struct {
	Root *Node
}

type Node struct {
	Value *Item
	Left  *Node
	Right *Node
}

func NewTreap() *Treap {
	return &Treap{
		Root: nil,
	}
}
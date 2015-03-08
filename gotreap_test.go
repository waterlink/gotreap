package gotreap

import "testing"

type MyItem struct {
	key      int
	priority int
	value    string
}

func (this *MyItem) LessKey(other *Item) bool {
	if otherItem, ok := (*other).(*MyItem); ok {
		return this.key < otherItem.key
	} else {
		panic("MyItem#LessKey :: Should have been a *MyItem")
	}
}

func (this *MyItem) LessPriority(other *Item) bool {
	if otherItem, ok := (*other).(*MyItem); ok {
		return this.priority < otherItem.priority
	} else {
		panic("MyItem#LessPriority :: Should have been a *MyItem")
	}
}

func TestCreatingNewTreap(t *testing.T) {
	NewTreap()
}

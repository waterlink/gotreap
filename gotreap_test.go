package gotreap

import "testing"

type MyItem struct {
	key      int
	priority int
	value    string
}

func (this *MyItem) LessKey(other Item) bool {
	if otherItem, ok := other.(*MyItem); ok {
		return this.key < otherItem.key
	} else {
		panic("MyItem#LessKey :: Should have been a *MyItem")
	}
}

func (this *MyItem) LessPriority(other Item) bool {
	if otherItem, ok := other.(*MyItem); ok {
		return this.priority < otherItem.priority
	} else {
		panic("MyItem#LessPriority :: Should have been a *MyItem")
	}
}

func TestCreatingNewTreap(t *testing.T) {
	NewTreap()
}

func expectToBeEq(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected to be equal to %s, but was %s", expected, actual)
	}
}

func expectToFind(t *testing.T, aSet *Treap, key int, expected string) {
	if foundItem := aSet.Search(&MyItem{key: key}); foundItem != nil {
		anItem, _ := foundItem.(*MyItem)
		expectToBeEq(t, anItem.value, expected)
	} else {
		t.Errorf("Unable to find item with key %d", key)
	}
}

func TestInsertSimple(t *testing.T) {
	myset := NewTreap()

	myset = myset.Insert(&MyItem{
		key:      5,
		priority: 17,
		value:    "hello world",
	})
	anItem := (*myset.Root.Value).(*MyItem)
	expectToBeEq(t, "hello world", anItem.value)

	myset = myset.Insert(&MyItem{
		key:      9,
		priority: 25,
		value:    "hedgehog",
	})
	anotherItem := (*myset.Root.Value).(*MyItem)
	expectToBeEq(t, "hedgehog", anotherItem.value)
}

func TestSeach(t *testing.T) {
	myset := NewTreap()

	myset = myset.Insert(&MyItem{
		key:      5,
		priority: 17,
		value:    "hello world",
	})
	expectToFind(t, myset, 5, "hello world")

	myset = myset.Insert(&MyItem{
		key:      9,
		priority: 25,
		value:    "hedgehog",
	})
	expectToFind(t, myset, 5, "hello world")
	expectToFind(t, myset, 9, "hedgehog")
}

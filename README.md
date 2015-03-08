# gotreap

Descartes Tree (treap, deramida) implementation in Golang.

## Installation

```go
import "github.com/waterlink/gotreap"
```

And run:

```bash
go get
```

## Usage

Define your custom treap (set up comparison rules for your items):

```go
type MyItem struct {
  key      int
  priority int
  value    string
}

func (this *MyItem) LessKey(other gotreap.Item) bool {
  if otherItem, ok := other.(*MyItem); ok {
    return this.key < otherItem.key
  } else {
    return false    // or panic here
  }
}

func (this *MyItem) LessPriority(other gotreap.Item) bool {
  if otherItem, ok := other.(*MyItem); ok {
    return this.priority < otherItem.priority
  } else {
    return false    // or panic here
  }
}
```

### Create new empty treap

Complexity: `O(1)`

```go
myset := gotreap.NewTreap()
```

### Insert an item into your treap

Complexity: `O(log N)`, where `N` - size of `myset`

```go
anItem := &MyItem{
  key:      5,
  priority: 17,
  value:    "hello world",
}
myset = myset.Insert(anItem)
```

### Search for an item in your treap

*Not implemented*

Complexity: `O(log N)`

```go
if foundItem := myset.Search(5); foundItem != nil {
  if anItem, ok := foundItem.(*MyItem); ok {
    // .. do something with anItem ..
  }
}
```

### Erase an item from your treap

*Not implemented*

Complexity: `O(log N)`

```go
myset = myset.Erase(5)
```

### Make a union of two treaps

*Not implemented*

Complexity: `O(M log (N/M))`, where `M` - size of `otherset`

```go
myset = gotreap.Union(myset, otherset)
```

## Contributing

1. Fork it ( https://github.com/waterlink/gotreap/fork )
2. Create a branch ( git checkout -b a-feature )
3. Commit your changes ( git commit -am "A feature" )
4. Push to the branch ( git push -u origin a-feature )
5. Create a new Pull Request

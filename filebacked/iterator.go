package filebacked

// Iterator interface.
// Read-only collection with stateful iteration.
type Iterator interface {
	// Len returns the number of items.
	Len() int
	// Reverse the content.
	Reverse()
	// At returns the object at index.
	At(index int) interface{}
	// AtWith populates with the objet at index.
	AtWith(int, interface{})
	// Next returns the next object.
	Next() (interface{}, bool)
	// NextWith returns the next object (with).
	NextWith(object interface{}) bool
	// Close the iterator.
	Close()
}

// FbIterator is a filebacked iterator.
type FbIterator struct {
	// Reader.
	*Reader
	// Current position.
	current int
}

// Next object.
func (r *FbIterator) Next() (object interface{}, hasNext bool) {
	if r.current < r.Len() {
		object = r.At(r.current)
		r.current++
		hasNext = true
	}

	return
}

// NextWith returns the next object.
func (r *FbIterator) NextWith(object interface{}) (hasNext bool) {
	if r.current < r.Len() {
		r.AtWith(r.current, object)
		r.current++
		hasNext = true
	}

	return
}

// Reverse the list.
func (r *FbIterator) Reverse() {
	in := r.index
	if len(in) == 0 {
		return
	}
	reversed := []int64{}
	for i := len(in) - 1; i >= 0; i-- {
		reversed = append(
			reversed,
			in[i])
	}

	r.index = reversed
}

// EmptyIterator is an empty iterator.
type EmptyIterator struct {
}

func (*EmptyIterator) Reverse() {
}

func (*EmptyIterator) Len() int {
	return 0
}

func (*EmptyIterator) At(int) interface{} {
	return nil
}

func (*EmptyIterator) AtWith(int, interface{}) {
	return
}

func (*EmptyIterator) Next() (interface{}, bool) {
	return nil, false
}

func (*EmptyIterator) NextWith(object interface{}) bool {
	return false
}

func (*EmptyIterator) Close() {
}

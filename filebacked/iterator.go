package filebacked

// Iterator interface.
// Read-only collection with stateful iteration.
type Iterator interface {
	// Len returns the number of items.
	Len() int
	// Reverse the content.
	Reverse()
	// At returns the object at index.
	At(index int) any
	// AtWith populates with the objet at index.
	AtWith(int, any)
	// Next returns the next object.
	Next() (any, bool)
	// NextWith returns the next object (with).
	NextWith(object any) bool
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
func (r *FbIterator) Next() (object any, hasNext bool) {
	if r.current < r.Len() {
		object = r.At(r.current)
		r.current++
		hasNext = true
	}

	return
}

// NextWith returns the next object.
func (r *FbIterator) NextWith(object any) (hasNext bool) {
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

func (*EmptyIterator) At(int) any {
	return nil
}

func (*EmptyIterator) AtWith(int, any) {
	return
}

func (*EmptyIterator) Next() (any, bool) {
	return nil, false
}

func (*EmptyIterator) NextWith(object any) bool {
	return false
}

func (*EmptyIterator) Close() {
}

package q

type Queue[T any] struct {
	start *QueueItem[T]
	end   *QueueItem[T]
	size  int
}

func NewQueue[T any]() (queue *Queue[T]) {
	queue = &Queue[T]{
		start: nil,
		end:   nil,
		size:  0,
	}

	return queue
}

func (q *Queue[T]) Push(value *T) {
	item := q.newItem(value)
	defer func() {
		q.size++
	}()

	if q.size != 0 {
		item.prev = q.end
		q.end.next = item
		q.end = item
		return
	}

	q.start = item
	q.end = item
	return
}

func (q *Queue[T]) Pop() (value *T) {
	if q.size > 1 {
		value = q.start.value

		q.start = q.start.next
		q.start.prev.next = nil
		q.start.prev = nil
		q.size--

		return value
	}

	if q.size == 1 {
		value = q.start.value

		q.start = nil
		q.end = nil
		q.size--

		return value
	}

	if q.size == 0 {
		return nil
	}

	panic("negative queue size")
}

func (q *Queue[T]) newItem(value *T) (qi *QueueItem[T]) {
	qi = &QueueItem[T]{
		queue: q,
		prev:  nil,
		next:  nil,
		value: value,
	}

	return qi
}

func (q *Queue[T]) Size() (size int) {
	return q.size
}

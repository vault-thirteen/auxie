package q

type QueueItem[T any] struct {
	queue *Queue[T]
	prev  *QueueItem[T]
	next  *QueueItem[T]
	value *T
}

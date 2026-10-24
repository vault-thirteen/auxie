package q

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewQueue(t *testing.T) {
	aTest := tester.New(t)
	queue := NewQueue[string]()
	aTest.MustBeEqual(queue.start, (*QueueItem[string])(nil))
	aTest.MustBeEqual(queue.end, (*QueueItem[string])(nil))
	aTest.MustBeEqual(queue.size, 0)
}

func Test_Push(t *testing.T) {
	type ClassA struct {
		Text string
	}

	aTest := tester.New(t)
	queue := NewQueue[ClassA]()

	// Step 1.
	var objA = ClassA{Text: "A"}
	{
		queue.Push(&objA)
		aTest.MustBeEqual(queue.size, 1)
		aTest.MustBeEqual(*queue.start.value, objA)
		aTest.MustBeEqual(queue.start, queue.end)

		aTest.MustBeEqual(queue.start.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(queue.start.next, (*QueueItem[ClassA])(nil))

		aTest.MustBeEqual(queue.end.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(queue.end.next, (*QueueItem[ClassA])(nil))
	}

	// Step 2.
	var objB = ClassA{Text: "B"}
	{
		queue.Push(&objB)
		aTest.MustBeEqual(queue.size, 2)
		aTest.MustBeEqual(*queue.start.value, objA)
		aTest.MustBeEqual(*queue.end.value, objB)

		aTest.MustBeEqual(queue.start.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(queue.start.next, queue.end)

		aTest.MustBeEqual(queue.end.prev, queue.start)
		aTest.MustBeEqual(queue.end.next, (*QueueItem[ClassA])(nil))
	}

	// Step 3.
	var objC = ClassA{Text: "C"}
	{
		queue.Push(&objC)
		aTest.MustBeEqual(queue.size, 3)
		aTest.MustBeEqual(*queue.start.value, objA)
		aTest.MustBeEqual(*queue.end.value, objC)

		aTest.MustBeEqual(queue.start.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(*queue.start.next.value, objB)

		aTest.MustBeEqual(*queue.start.next.prev.value, objA)
		aTest.MustBeEqual(*queue.start.next.next.value, objC)

		aTest.MustBeEqual(*queue.end.prev.value, objB)
		aTest.MustBeEqual(queue.end.next, (*QueueItem[ClassA])(nil))

		aTest.MustBeEqual(*queue.end.prev.prev.value, objA)
		aTest.MustBeEqual(*queue.end.prev.next.value, objC)
	}
}

func Test_Pop(t *testing.T) {
	type ClassA struct {
		Text string
	}

	aTest := tester.New(t)
	queue := NewQueue[ClassA]()
	var objA = ClassA{Text: "A"}
	var objB = ClassA{Text: "B"}
	var objC = ClassA{Text: "C"}
	queue.Push(&objA)
	queue.Push(&objB)
	queue.Push(&objC)

	var x *ClassA

	// Step 1.
	{
		x = queue.Pop()
		aTest.MustBeEqual(*x, objA)
		aTest.MustBeEqual(queue.size, 2)

		aTest.MustBeEqual(*queue.start.value, objB)
		aTest.MustBeEqual(*queue.end.value, objC)

		aTest.MustBeEqual(queue.start.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(queue.start.next, queue.end)

		aTest.MustBeEqual(queue.end.prev, queue.start)
		aTest.MustBeEqual(queue.end.next, (*QueueItem[ClassA])(nil))
	}

	// Step 2.
	{
		x = queue.Pop()
		aTest.MustBeEqual(*x, objB)
		aTest.MustBeEqual(queue.size, 1)

		aTest.MustBeEqual(*queue.start.value, objC)
		aTest.MustBeEqual(*queue.end.value, objC)

		aTest.MustBeEqual(queue.start, queue.end)
		aTest.MustBeEqual(queue.start.prev, (*QueueItem[ClassA])(nil))
		aTest.MustBeEqual(queue.start.next, (*QueueItem[ClassA])(nil))
	}

	// Step 3.
	{
		x = queue.Pop()
		aTest.MustBeEqual(*x, objC)
		aTest.MustBeEqual(queue.size, 0)

		aTest.MustBeEqual(queue.start, queue.end)
		aTest.MustBeEqual(queue.start, (*QueueItem[ClassA])(nil))
	}

	// Step 4.
	{
		x = queue.Pop()
		aTest.MustBeEqual(x, (*ClassA)(nil))
		aTest.MustBeEqual(queue.size, 0)
	}

	// Step 5.
	{
		queue = &Queue[ClassA]{size: -1}

		defer func() {
			rec := recover()
			aTest.MustBeEqual(rec, "negative queue size")
		}()
		_ = queue.Pop()
	}
}

func Test_Size(t *testing.T) {
	aTest := tester.New(t)
	queue := NewQueue[string]()
	aTest.MustBeEqual(queue.Size(), 0)

	sA := "A"
	queue.Push(&sA)
	aTest.MustBeEqual(queue.Size(), 1)

	sB := "B"
	queue.Push(&sB)
	aTest.MustBeEqual(queue.Size(), 2)

	_ = queue.Pop()
	aTest.MustBeEqual(queue.Size(), 1)

	_ = queue.Pop()
	aTest.MustBeEqual(queue.Size(), 0)

	_ = queue.Pop()
	aTest.MustBeEqual(queue.Size(), 0)
}

func Test_newItem(t *testing.T) {
	aTest := tester.New(t)
	queue := &Queue[int]{}

	i := 1
	x := queue.newItem(&i)
	aTest.MustBeEqual(x.queue, queue)
	aTest.MustBeEqual(x.prev, (*QueueItem[int])(nil))
	aTest.MustBeEqual(x.next, (*QueueItem[int])(nil))
	aTest.MustBeEqual(x.value, &i)
}

package queue

import "testing"

func TestAddQueue(t *testing.T) {
	// empty queue with capacity of 3
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf(
				"incorrect queue items count:\n\tExpected: %d\n\tGot: %d",
				i,
				len(q.items),
			)
		}

		if !q.Append(i) {
			t.Errorf("failed to append item %d to queue", i)
		}
	}

	// now capacity if full (3)
	if q.Append(4) {
		t.Errorf("should not add an item to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}

		if item != i {
			t.Errorf(
				"should get items in right order:\n\tExpected: %d\n\tGot: %d",
				i, item,
			)
		}
	}

	// queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("should not get any items from an empty queue, got: %d", item)
	}
}

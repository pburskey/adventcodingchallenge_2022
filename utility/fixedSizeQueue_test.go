package utility

import (
	"testing"
)

func TestNewFixedSizeStringQueue(t *testing.T) {
	queue := NewFixedSizeStringQueue(4)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")
	duplicates := FindDuplicates(queue.PeakAll())
	if duplicates != nil {

	}

	queue = NewFixedSizeStringQueue(4)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	duplicates = FindDuplicates(queue.PeakAll())
	if duplicates != nil {

	}

	queue = NewFixedSizeStringQueue(4)
	queue.Enqueue("a")
	queue.Enqueue("a")
	queue.Enqueue("a")
	queue.Enqueue("a")

	duplicates = FindDuplicates(queue.PeakAll())
	if duplicates != nil {

	}

}

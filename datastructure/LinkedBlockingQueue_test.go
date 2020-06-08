package TestProject

import "testing"

func TestLinkedBlockingQueue(t *testing.T) {
	lbq := NewLInkedBlockingQueue(10)
	t.Log(lbq)
}

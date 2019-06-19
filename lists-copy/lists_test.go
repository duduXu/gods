package arraylist_copy

import (
	"testing"
)

func TestLists(t *testing.T)  {
	list := New()

	list.Add(1, 3, 89, 5, 6, 37, 58)
	t.Log(list)

	//list.Remove(3)
	//t.Log(list)
}

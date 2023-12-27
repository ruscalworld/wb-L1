package task7

import "testing"

func TestSyncMap(t *testing.T) {
	m := MakeMap[string, int](2)

	_, ok := m.Load("123")
	if ok {
		t.Errorf("Value for key 123 should be considered as unexistent")
		return
	}

	m.Store("1", 12)
	v, ok := m.Load("1")
	if !ok || v != 12 {
		t.Errorf("Value for key 1 should be considered as existing and be equal to 12")
		return
	}
}

package prototype

import "testing"

func TestExample_Clone(t *testing.T) {
	origin := NewServer()
	current := origin.Clone()
	t.Log(origin)
	t.Log(current)
}

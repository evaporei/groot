package cli

import "testing"

func TestAreArgsFilled(t *testing.T) {
	t.Run("areArgsFilled([]) == false", func(t *testing.T) {
		args := make([]string, 0)
		if areArgsFilled(args) != false {
			t.Fail()
		}
	})
	t.Run("areArgsFilled([\"a\",\"b\"]) == true", func(t *testing.T) {
		args := []string{"a", "b"}
		if areArgsFilled(args) != true {
			t.Fail()
		}
	})
}
